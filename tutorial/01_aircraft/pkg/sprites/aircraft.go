package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Aircraft struct {
	Sprite
	timer      float32
	moveSpeed  float32
	OnDieEvent *Event0
}

func (pself *Aircraft) OnStart() {
	pself.moveSpeed = 600
	pself.OnDieEvent = NewEvent0()
}
func (pself *Aircraft) OnUpdate(delta float32) {
	pself.timer += delta
	if pself.timer > 0.3 {
		pself.timer = 0
		bullet := CreateSprite[Bullet]()
		pos := pself.GetPosition()
		bullet.SetPosition(Vec2{pos.X, pos.Y + 150})
	}

	if InputMgr.GetKey(KeyCode.W) {
		pself.Move(0, pself.moveSpeed*delta)
	}
	if InputMgr.GetKey(KeyCode.S) {
		pself.Move(0, -pself.moveSpeed*delta)
	}
	if InputMgr.GetKey(KeyCode.D) {
		pself.Move(pself.moveSpeed*delta, 0)
	}
	if InputMgr.GetKey(KeyCode.A) {
		pself.Move(-pself.moveSpeed*delta, 0)
	}
}

func (pself *Aircraft) OnDestory() {
	println("Aircraft OnDestory")
}

func (pself *Aircraft) OnHit() {
	pself.OnDieEvent.Trigger()
}