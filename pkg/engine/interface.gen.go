/*
------------------------------------------------------------------------------
//   This code was generated by template ffi_gdextension_interface.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_gdextension_interface.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package engine

var (
	AudioMgr    IAudioMgr
	CameraMgr   ICameraMgr
	InputMgr    IInputMgr
	PhysicMgr   IPhysicMgr
	PlatformMgr IPlatformMgr
	ResMgr      IResMgr
	SceneMgr    ISceneMgr
	SpriteMgr   ISpriteMgr
	UiMgr       IUiMgr
)

type IAudioMgr interface {
	StopAll()
	PlaySfx(path string)
	PlayMusic(path string)
	PauseMusic()
	ResumeMusic()
	GetMusicTimer() float32
	SetMusicTimer(time float32)
	IsMusicPlaying() bool
	SetSfxVolume(volume float32)
	GetSfxVolume() float32
	SetMusicVolume(volume float32)
	GetMusicVolume() float32
	SetMasterVolume(volume float32)
	GetMasterVolume() float32
}

type ICameraMgr interface {
	GetCameraPosition() Vec2
	SetCameraPosition(position Vec2)
	GetCameraZoom() Vec2
	SetCameraZoom(size Vec2)
	GetViewportRect() Rect2
}

type IInputMgr interface {
	GetMousePos() Vec2
	GetKey(key int64) bool
	GetMouseState(mouse_id int64) bool
	GetKeyState(key int64) int64
	GetAxis(neg_action string, pos_action string) float32
	IsActionPressed(action string) bool
	IsActionJustPressed(action string) bool
	IsActionJustReleased(action string) bool
}

type IPhysicMgr interface {
	Raycast(from Vec2, to Vec2, collision_mask int64) Object
	CheckCollision(from Vec2, to Vec2, collision_mask int64, collide_with_areas bool, collide_with_bodies bool) bool
}

type IPlatformMgr interface {
	SetWindowSize(width int64, height int64)
	GetWindowSize() Vec2
	SetWindowTitle(title string)
	GetWindowTitle() string
	SetWindowFullscreen(enable bool)
	IsWindowFullscreen() bool
	SetDebugMode(enable bool)
	IsDebugMode() bool
}

type IResMgr interface {
	GetImageSize(path string) Vec2
}

type ISceneMgr interface {
	ChangeSceneToFile(path string)
	ReloadCurrentScene() int64
	UnloadCurrentScene()
}

type ISpriteMgr interface {
	SetDontDestroyOnLoad(obj Object)
	SetProcess(obj Object, is_on bool)
	SetPhysicProcess(obj Object, is_on bool)
	SetChildPosition(obj Object, path string, pos Vec2)
	GetChildPosition(obj Object, path string) Vec2
	SetChildRotation(obj Object, path string, rot float32)
	GetChildRotation(obj Object, path string) float32
	SetChildScale(obj Object, path string, scale Vec2)
	GetChildScale(obj Object, path string) Vec2
	CheckCollision(obj Object, target Object, is_src_trigger bool, is_dst_trigger bool) bool
	CheckCollisionWithPoint(obj Object, point Vec2, is_trigger bool) bool
	CreateSprite(path string) Object
	CloneSprite(obj Object) Object
	DestroySprite(obj Object) bool
	IsSpriteAlive(obj Object) bool
	SetPosition(obj Object, pos Vec2)
	GetPosition(obj Object) Vec2
	SetRotation(obj Object, rot float32)
	GetRotation(obj Object) float32
	SetScale(obj Object, scale Vec2)
	GetScale(obj Object) Vec2
	SetColor(obj Object, color Color)
	GetColor(obj Object) Color
	SetTextureAltas(obj Object, path string, rect2 Rect2)
	SetTexture(obj Object, path string)
	GetTexture(obj Object) string
	SetVisible(obj Object, visible bool)
	GetVisible(obj Object) bool
	GetZIndex(obj Object) int64
	SetZIndex(obj Object, z int64)
	PlayAnim(obj Object, p_name string, p_custom_scale float32, p_from_end bool)
	PlayBackwardsAnim(obj Object, p_name string)
	PauseAnim(obj Object)
	StopAnim(obj Object)
	IsPlayingAnim(obj Object) bool
	SetAnim(obj Object, p_name string)
	GetAnim(obj Object) string
	SetAnimFrame(obj Object, p_frame int64)
	GetAnimFrame(obj Object) int64
	SetAnimSpeedScale(obj Object, p_speed_scale float32)
	GetAnimSpeedScale(obj Object) float32
	GetAnimPlayingSpeed(obj Object) float32
	SetAnimCentered(obj Object, p_center bool)
	IsAnimCentered(obj Object) bool
	SetAnimOffset(obj Object, p_offset Vec2)
	GetAnimOffset(obj Object) Vec2
	SetAnimFlipH(obj Object, p_flip bool)
	IsAnimFlippedH(obj Object) bool
	SetAnimFlipV(obj Object, p_flip bool)
	IsAnimFlippedV(obj Object) bool
	SetVelocity(obj Object, velocity Vec2)
	GetVelocity(obj Object) Vec2
	IsOnFloor(obj Object) bool
	IsOnFloorOnly(obj Object) bool
	IsOnWall(obj Object) bool
	IsOnWallOnly(obj Object) bool
	IsOnCeiling(obj Object) bool
	IsOnCeilingOnly(obj Object) bool
	GetLastMotion(obj Object) Vec2
	GetPositionDelta(obj Object) Vec2
	GetFloorNormal(obj Object) Vec2
	GetWallNormal(obj Object) Vec2
	GetRealVelocity(obj Object) Vec2
	MoveAndSlide(obj Object)
	SetGravity(obj Object, gravity float32)
	GetGravity(obj Object) float32
	SetMass(obj Object, mass float32)
	GetMass(obj Object) float32
	AddForce(obj Object, force Vec2)
	AddImpulse(obj Object, impulse Vec2)
	SetCollisionLayer(obj Object, layer int64)
	GetCollisionLayer(obj Object) int64
	SetCollisionMask(obj Object, mask int64)
	GetCollisionMask(obj Object) int64
	SetTriggerLayer(obj Object, layer int64)
	GetTriggerLayer(obj Object) int64
	SetTriggerMask(obj Object, mask int64)
	GetTriggerMask(obj Object) int64
	SetColliderRect(obj Object, center Vec2, size Vec2)
	SetColliderCircle(obj Object, center Vec2, radius float32)
	SetColliderCapsule(obj Object, center Vec2, size Vec2)
	SetCollisionEnabled(obj Object, enabled bool)
	IsCollisionEnabled(obj Object) bool
	SetTriggerRect(obj Object, center Vec2, size Vec2)
	SetTriggerCircle(obj Object, center Vec2, radius float32)
	SetTriggerCapsule(obj Object, center Vec2, size Vec2)
	SetTriggerEnabled(obj Object, trigger bool)
	IsTriggerEnabled(obj Object) bool
}

type IUiMgr interface {
	BindNode(obj Object, rel_path string) Object
	CreateNode(path string) Object
	CreateButton(path string, text string) Object
	CreateLabel(path string, text string) Object
	CreateImage(path string) Object
	CreateToggle(path string, value bool) Object
	CreateSlider(path string, value float32) Object
	CreateInput(path string, text string) Object
	DestroyNode(obj Object) bool
	GetType(obj Object) int64
	SetText(obj Object, text string)
	GetText(obj Object) string
	SetTexture(obj Object, path string)
	GetTexture(obj Object) string
	SetColor(obj Object, color Color)
	GetColor(obj Object) Color
	SetFontSize(obj Object, size int64)
	GetFontSize(obj Object) int64
	SetVisible(obj Object, visible bool)
	GetVisible(obj Object) bool
	SetInteractable(obj Object, interactable bool)
	GetInteractable(obj Object) bool
	SetRect(obj Object, rect Rect2)
	GetRect(obj Object) Rect2
	GetLayoutDirection(obj Object) int64
	SetLayoutDirection(obj Object, value int64)
	GetLayoutMode(obj Object) int64
	SetLayoutMode(obj Object, value int64)
	GetAnchorsPreset(obj Object) int64
	SetAnchorsPreset(obj Object, value int64)
	GetScale(obj Object) Vec2
	SetScale(obj Object, value Vec2)
	GetPosition(obj Object) Vec2
	SetPosition(obj Object, value Vec2)
	GetSize(obj Object) Vec2
	SetSize(obj Object, value Vec2)
	GetGlobalPosition(obj Object) Vec2
	SetGlobalPosition(obj Object, value Vec2)
	GetRotation(obj Object) float32
	SetRotation(obj Object, value float32)
	GetFlip(obj Object, horizontal bool) bool
	SetFlip(obj Object, horizontal bool, is_flip bool)
}
