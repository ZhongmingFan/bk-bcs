/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package manager

import (
	"context"
	"strconv"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/go-redis/redis/v8"
	"go-micro.dev/v4/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PodCleanUpManager Pod 清理
type PodCleanUpManager struct {
	ctx         context.Context
	redisClient *redis.Client
	k8sClient   *kubernetes.Clientset
}

func NewPodCleanUpManager(ctx context.Context) *PodCleanUpManager {
	return &PodCleanUpManager{
		ctx: ctx,
	}
}

// 记录pod心跳
// 定时上报存活, 清理时需要使用
func (p *PodCleanUpManager) Heartbeat(podName string) {
	now := float64(time.Now().Unix())
	p.redisClient.ZAdd(p.ctx, WebConsoleHeartbeatKey, &redis.Z{Member: podName, Score: now})
}

// getActiveUserPod 获取存活节点
func (p *PodCleanUpManager) getActiveUserPod() []string {
	startTime := time.Now().Add(-UserPodExpireTime).Format("20060102150405")
	// 删除掉过期数据
	p.redisClient.ZRemRangeByScore(p.ctx, WebConsoleHeartbeatKey, "-inf", startTime)

	// 获取存活的pod
	activatedPods := p.redisClient.ZRange(p.ctx, WebConsoleHeartbeatKey, 0, -1).Val()

	return activatedPods
}

// CleanUserPod 单个集群清理
func (p *PodCleanUpManager) CleanUserPod() error {

	// TODO 根据不同的集群进行删除

	alivePods := p.getActiveUserPod()
	alivePodsMap := make(map[string]string)
	for _, pod := range alivePods {
		alivePodsMap[pod] = pod
	}

	podList, err := p.k8sClient.CoreV1().Pods(Namespace).List(p.ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	p.cleanUserPodByCluster(podList, alivePodsMap)
	return nil

}

func (p *PodCleanUpManager) Run() error {
	interval := time.NewTicker(10 * time.Second)
	defer interval.Stop()

	for {
		select {
		case <-p.ctx.Done():
			logger.Info("close PodCleanUpManager done")
			return p.ctx.Err()
		case <-interval.C:
			if err := p.CleanUserPod(); err != nil {
				logger.Errorf("clean use pod error, %s", err)
			}
		}
	}
}

// 清理用户下的相关集群pod
func (p *PodCleanUpManager) cleanUserPodByCluster(podList *v1.PodList, alivePods map[string]string) {

	// 过期时间
	timeDiff, _ := time.ParseDuration("-" + strconv.FormatInt(UserPodExpireTime, 10) + "s")
	minExpireTime := time.Now().Add(timeDiff) // 在此时间之前的都算作过期

	for _, pod := range podList.Items {
		if pod.Status.Phase == "Pending" {
			continue
		}

		// 小于一个周期的pod不清理
		podCreateTimeStr, ok := pod.ObjectMeta.Labels[LabelWebConsoleCreateTimestamp]
		if !ok {
			continue
		}

		podCreateTime, _ := time.Parse("20060102150405", podCreateTimeStr)
		if minExpireTime.After(podCreateTime) {
			blog.Info("pod %s exist time %s > %s, just ignore", pod.Name, podCreateTimeStr, minExpireTime)
			continue
		}

		// 有心跳上报的不清理
		if _, ok := alivePods[pod.Name]; ok {
			continue
		}

		// 删除pod
		err := p.k8sClient.CoreV1().Pods(Namespace).Delete(p.ctx, pod.Name, metav1.DeleteOptions{})
		if err != nil {
			blog.Errorf("delete pod(%s) failed, err: %v", pod.Name, err)
			continue
		}
		blog.Info("delete pod %s", pod.Name)

		// 删除configMap
		for _, volume := range pod.Spec.Volumes {
			if volume.ConfigMap != nil {
				err = p.k8sClient.CoreV1().ConfigMaps(Namespace).Delete(p.ctx,
					volume.ConfigMap.LocalObjectReference.Name, metav1.DeleteOptions{})
				if err != nil {
					blog.Errorf("delete configmap %s failed ,err : %v", volume.ConfigMap.LocalObjectReference.Name, err)
				}
				blog.Info("delete configmap %s", volume.ConfigMap.LocalObjectReference.Name)
			}
		}
	}
}