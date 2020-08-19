package golib

import (
	"time"
)


//计时器，方法开始时执行一次；然后每天晚上12点执行
func StartTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

var retryTime = 0

//根据重试次数，来计算下次重试的时间（上限30min）
func getRetryDuration() time.Duration {
	switch retryTime {
	case 0:
		return 10
	case 1:
		return 180
	case 2:
		return 600
	case 3:
		return 1800
	default:
		return 1800
	}
}

//以下为定时任务如何使用的示例

////开始时执行一次，然后每到晚上12点执行一次f
//func StartTimerPerMidNight(f func() error) {
//	go func() {
//		for {
//			now := time.Now()
//			var t *time.Timer
//			logger.Info("start handle expire apply and approval")
//			err := f()
//			如果未发生错误，第二天0点重试
//			if err == nil {
//				retryTime = 0
//				// 计算下一个零点
//				next := now.Add(time.Hour * 24)
//				next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
//				logger.Info("success handle expire apply and approval, next handle time :%v", next.String())
//				t = time.NewTimer(next.Sub(now))
//			如果发生错误，根据指定的时间周期重试
//			} else {
//				retryTime++
//				logger.Errorf("handle expire apply and approval error:%v, and after %v seconds, start %vth retry",
//					err, time.Second*getRetryDuration(), retryTime)
//				next := now.Add(time.Second * getRetryDuration())
//				next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
//				t = time.NewTimer(next.Sub(now))
//			}
//			<-t.C
//		}
//	}()
//}
//
//func HandleExpireApplyAndApproval() error {
//
//	currentTime := time.Now().String()
//	applyExpireState := []uint8{
//		model.APPLY_STATE_NEED_ORG_ADMIN_APPROVAL,
//		model.APPLY_STATE_APPROVING,
//	}
//	//获取所有过期的申请
//	expireApplyIds := []uint{}
//
//	err := config.GetDBIns().Model(&model.Apply{}).
//		Where(" expire_at < ? AND state IN (?) ", currentTime, tools.ChangeUint8ArrayToString(applyExpireState)).
//		Pluck("id", &expireApplyIds).Error
//	if err != nil {
//		return err
//	}
//	if len(expireApplyIds) == 0 {
//		return errors.New("fuck fuck")
//	}
//	//更新过期申请的状态
//	tx := config.GetDBIns().Begin()
//	err = tx.Table("apply").
//		Where(" id IN (?) ", tools.ChangeUintArrayToString(expireApplyIds)).
//		Updates(map[string]interface{}{"state": model.APPLY_STATE_EXPIRED}).Error
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//	//更新过期申请关联审批的状态
//	err = tx.Table("approval").
//		Where(" apply_id IN (?) ", tools.ChangeUintArrayToString(expireApplyIds)).
//		Updates(map[string]interface{}{"state": model.APPROVAL_STATE_EXPIRED}).Error
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//
//	return nil
//}