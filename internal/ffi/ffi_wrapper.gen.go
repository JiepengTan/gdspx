/*------------------------------------------------------------------------------
//   This code was generated by template ffi_wrapper.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_wrapper.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/
package ffi

//revive:disable

// #include "gdextension_spx_codegen_header.h"
// #include "ffi_wrapper.gen.h"
// #include <stdint.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
)

// C type aliases
// enums


// C function aliases
type GDExtensionSpxGlobalRegisterCallbacks C.GDExtensionSpxGlobalRegisterCallbacks
type GDExtensionSpxInputGetMousePos C.GDExtensionSpxInputGetMousePos
type GDExtensionSpxInputGetMouseState C.GDExtensionSpxInputGetMouseState
type GDExtensionSpxInputGetKeyState C.GDExtensionSpxInputGetKeyState
type GDExtensionSpxInputGetAxis C.GDExtensionSpxInputGetAxis
type GDExtensionSpxInputIsActionPressed C.GDExtensionSpxInputIsActionPressed
type GDExtensionSpxInputIsActionJustPressed C.GDExtensionSpxInputIsActionJustPressed
type GDExtensionSpxInputIsActionJustReleased C.GDExtensionSpxInputIsActionJustReleased
type GDExtensionSpxAudioPlayAudio C.GDExtensionSpxAudioPlayAudio
type GDExtensionSpxAudioSetAudioVolume C.GDExtensionSpxAudioSetAudioVolume
type GDExtensionSpxAudioGetAudioVolume C.GDExtensionSpxAudioGetAudioVolume
type GDExtensionSpxAudioIsMusicPlaying C.GDExtensionSpxAudioIsMusicPlaying
type GDExtensionSpxAudioPlayMusic C.GDExtensionSpxAudioPlayMusic
type GDExtensionSpxAudioSetMusicVolume C.GDExtensionSpxAudioSetMusicVolume
type GDExtensionSpxAudioGetMusicVolume C.GDExtensionSpxAudioGetMusicVolume
type GDExtensionSpxAudioPauseMusic C.GDExtensionSpxAudioPauseMusic
type GDExtensionSpxAudioResumeMusic C.GDExtensionSpxAudioResumeMusic
type GDExtensionSpxAudioGetMusicTimer C.GDExtensionSpxAudioGetMusicTimer
type GDExtensionSpxAudioSetMusicTimer C.GDExtensionSpxAudioSetMusicTimer
type GDExtensionSpxPhysicSetGravity C.GDExtensionSpxPhysicSetGravity
type GDExtensionSpxPhysicGetGravity C.GDExtensionSpxPhysicGetGravity
type GDExtensionSpxPhysicSetVelocity C.GDExtensionSpxPhysicSetVelocity
type GDExtensionSpxPhysicGetVelocity C.GDExtensionSpxPhysicGetVelocity
type GDExtensionSpxPhysicSetMass C.GDExtensionSpxPhysicSetMass
type GDExtensionSpxPhysicGetMass C.GDExtensionSpxPhysicGetMass
type GDExtensionSpxPhysicAddForce C.GDExtensionSpxPhysicAddForce
type GDExtensionSpxPhysicAddImpulse C.GDExtensionSpxPhysicAddImpulse
type GDExtensionSpxPhysicSetCollisionLayer C.GDExtensionSpxPhysicSetCollisionLayer
type GDExtensionSpxPhysicGetCollisionLayer C.GDExtensionSpxPhysicGetCollisionLayer
type GDExtensionSpxPhysicSetCollisionMask C.GDExtensionSpxPhysicSetCollisionMask
type GDExtensionSpxPhysicGetCollisionMask C.GDExtensionSpxPhysicGetCollisionMask
type GDExtensionSpxPhysicGetColliderType C.GDExtensionSpxPhysicGetColliderType
type GDExtensionSpxPhysicAddColliderRect C.GDExtensionSpxPhysicAddColliderRect
type GDExtensionSpxPhysicAddColliderCircle C.GDExtensionSpxPhysicAddColliderCircle
type GDExtensionSpxPhysicAddColliderCapsule C.GDExtensionSpxPhysicAddColliderCapsule
type GDExtensionSpxPhysicSetTrigger C.GDExtensionSpxPhysicSetTrigger
type GDExtensionSpxPhysicIsTrigger C.GDExtensionSpxPhysicIsTrigger
type GDExtensionSpxPhysicSetCollisionEnabled C.GDExtensionSpxPhysicSetCollisionEnabled
type GDExtensionSpxPhysicIsCollisionEnabled C.GDExtensionSpxPhysicIsCollisionEnabled
type GDExtensionSpxSpriteCreateSprite C.GDExtensionSpxSpriteCreateSprite
type GDExtensionSpxSpriteCloneSprite C.GDExtensionSpxSpriteCloneSprite
type GDExtensionSpxSpriteDestroySprite C.GDExtensionSpxSpriteDestroySprite
type GDExtensionSpxSpriteIsSpriteAlive C.GDExtensionSpxSpriteIsSpriteAlive
type GDExtensionSpxSpriteSetPosition C.GDExtensionSpxSpriteSetPosition
type GDExtensionSpxSpriteSetRotation C.GDExtensionSpxSpriteSetRotation
type GDExtensionSpxSpriteSetScale C.GDExtensionSpxSpriteSetScale
type GDExtensionSpxSpriteGetPosition C.GDExtensionSpxSpriteGetPosition
type GDExtensionSpxSpriteGetRotation C.GDExtensionSpxSpriteGetRotation
type GDExtensionSpxSpriteGetScale C.GDExtensionSpxSpriteGetScale
type GDExtensionSpxSpriteSetColor C.GDExtensionSpxSpriteSetColor
type GDExtensionSpxSpriteGetColor C.GDExtensionSpxSpriteGetColor
type GDExtensionSpxSpriteUpdateTexture C.GDExtensionSpxSpriteUpdateTexture
type GDExtensionSpxSpriteGetTexture C.GDExtensionSpxSpriteGetTexture
type GDExtensionSpxSpriteSetVisible C.GDExtensionSpxSpriteSetVisible
type GDExtensionSpxSpriteGetVisible C.GDExtensionSpxSpriteGetVisible
type GDExtensionSpxSpriteUpdateZIndex C.GDExtensionSpxSpriteUpdateZIndex
type GDExtensionSpxUICreateButton C.GDExtensionSpxUICreateButton
type GDExtensionSpxUICreateLabel C.GDExtensionSpxUICreateLabel
type GDExtensionSpxUICreateImage C.GDExtensionSpxUICreateImage
type GDExtensionSpxUICreateSlider C.GDExtensionSpxUICreateSlider
type GDExtensionSpxUICreateToggle C.GDExtensionSpxUICreateToggle
type GDExtensionSpxUICreateInput C.GDExtensionSpxUICreateInput
type GDExtensionSpxUIGetType C.GDExtensionSpxUIGetType
type GDExtensionSpxUISetInteractable C.GDExtensionSpxUISetInteractable
type GDExtensionSpxUIGetInteractable C.GDExtensionSpxUIGetInteractable
type GDExtensionSpxUISetText C.GDExtensionSpxUISetText
type GDExtensionSpxUIGetText C.GDExtensionSpxUIGetText
type GDExtensionSpxUISetRect C.GDExtensionSpxUISetRect
type GDExtensionSpxUIGetRect C.GDExtensionSpxUIGetRect
type GDExtensionSpxUISetColor C.GDExtensionSpxUISetColor
type GDExtensionSpxUIGetColor C.GDExtensionSpxUIGetColor
type GDExtensionSpxUISetFontSize C.GDExtensionSpxUISetFontSize
type GDExtensionSpxUIGetFontSize C.GDExtensionSpxUIGetFontSize
type GDExtensionSpxUISetVisible C.GDExtensionSpxUISetVisible
type GDExtensionSpxUIGetVisible C.GDExtensionSpxUIGetVisible
type GDExtensionSpxCallbackOnEngineStart C.GDExtensionSpxCallbackOnEngineStart
type GDExtensionSpxCallbackOnEngineUpdate C.GDExtensionSpxCallbackOnEngineUpdate
type GDExtensionSpxCallbackOnEngineDestroy C.GDExtensionSpxCallbackOnEngineDestroy
type GDExtensionSpxCallbackOnSpriteReady C.GDExtensionSpxCallbackOnSpriteReady
type GDExtensionSpxCallbackOnSpriteUpdated C.GDExtensionSpxCallbackOnSpriteUpdated
type GDExtensionSpxCallbackOnSpriteDestroyed C.GDExtensionSpxCallbackOnSpriteDestroyed
type GDExtensionSpxCallbackOnMousePressed C.GDExtensionSpxCallbackOnMousePressed
type GDExtensionSpxCallbackOnMouseReleased C.GDExtensionSpxCallbackOnMouseReleased
type GDExtensionSpxCallbackOnKeyPressed C.GDExtensionSpxCallbackOnKeyPressed
type GDExtensionSpxCallbackOnKeyReleased C.GDExtensionSpxCallbackOnKeyReleased
type GDExtensionSpxCallbackOnActionPressed C.GDExtensionSpxCallbackOnActionPressed
type GDExtensionSpxCallbackOnActionJustPressed C.GDExtensionSpxCallbackOnActionJustPressed
type GDExtensionSpxCallbackOnActionJustReleased C.GDExtensionSpxCallbackOnActionJustReleased
type GDExtensionSpxCallbackOnAxisChanged C.GDExtensionSpxCallbackOnAxisChanged
type GDExtensionSpxCallbackOnCollisionEnter C.GDExtensionSpxCallbackOnCollisionEnter
type GDExtensionSpxCallbackOnCollisionStay C.GDExtensionSpxCallbackOnCollisionStay
type GDExtensionSpxCallbackOnCollisionExit C.GDExtensionSpxCallbackOnCollisionExit
type GDExtensionSpxCallbackOnTriggerEnter C.GDExtensionSpxCallbackOnTriggerEnter
type GDExtensionSpxCallbackOnTriggerStay C.GDExtensionSpxCallbackOnTriggerStay
type GDExtensionSpxCallbackOnTriggerExit C.GDExtensionSpxCallbackOnTriggerExit
type GDExtensionSpxCallbackOnUIPressed C.GDExtensionSpxCallbackOnUIPressed
type GDExtensionSpxCallbackOnUIReleased C.GDExtensionSpxCallbackOnUIReleased
type GDExtensionSpxCallbackOnUIHovered C.GDExtensionSpxCallbackOnUIHovered
type GDExtensionSpxCallbackOnUIClicked C.GDExtensionSpxCallbackOnUIClicked
type GDExtensionSpxCallbackOnUIToggle C.GDExtensionSpxCallbackOnUIToggle
type GDExtensionSpxCallbackOnUITextChanged C.GDExtensionSpxCallbackOnUITextChanged


// call gdextension interface functions
func CallFunc_GDExtensionSpxGlobalRegisterCallbacks(
	callback_ptr GDExtensionSpxCallbackInfoPtr,
	)  {
	arg0 := (C.GDExtensionSpxGlobalRegisterCallbacks)(FFI.SpxGlobalRegisterCallbacks)
	arg1 := (C.GDExtensionSpxCallbackInfoPtr)(callback_ptr)
	C.cgo_callfn_GDExtensionSpxGlobalRegisterCallbacks(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxInputGetMousePos(
	) GdVec2 {
	arg0 := (C.GDExtensionSpxInputGetMousePos)(FFI.SpxInputGetMousePos)
	ret := C.cgo_callfn_GDExtensionSpxInputGetMousePos(arg0,)
	return (GdVec2)(ret)
}
func CallFunc_GDExtensionSpxInputGetMouseState(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputGetMouseState)(FFI.SpxInputGetMouseState)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxInputGetMouseState(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxInputGetKeyState(
	key GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxInputGetKeyState)(FFI.SpxInputGetKeyState)
	arg1 := (C.GdInt)(key)
	ret := C.cgo_callfn_GDExtensionSpxInputGetKeyState(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxInputGetAxis(
	axis GdString,
	) GdFloat {
	arg0 := (C.GDExtensionSpxInputGetAxis)(FFI.SpxInputGetAxis)
	arg1 := (C.GdString)(axis)
	ret := C.cgo_callfn_GDExtensionSpxInputGetAxis(arg0,arg1,)
	
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxInputIsActionPressed(
	action GdString,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputIsActionPressed)(FFI.SpxInputIsActionPressed)
	arg1 := (C.GdString)(action)
	ret := C.cgo_callfn_GDExtensionSpxInputIsActionPressed(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxInputIsActionJustPressed(
	action GdString,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputIsActionJustPressed)(FFI.SpxInputIsActionJustPressed)
	arg1 := (C.GdString)(action)
	ret := C.cgo_callfn_GDExtensionSpxInputIsActionJustPressed(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxInputIsActionJustReleased(
	action GdString,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputIsActionJustReleased)(FFI.SpxInputIsActionJustReleased)
	arg1 := (C.GdString)(action)
	ret := C.cgo_callfn_GDExtensionSpxInputIsActionJustReleased(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxAudioPlayAudio(
	path GdString,
	)  {
	arg0 := (C.GDExtensionSpxAudioPlayAudio)(FFI.SpxAudioPlayAudio)
	arg1 := (C.GdString)(path)
	C.cgo_callfn_GDExtensionSpxAudioPlayAudio(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxAudioSetAudioVolume(
	volume GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxAudioSetAudioVolume)(FFI.SpxAudioSetAudioVolume)
	arg1 := (C.GdFloat)(volume)
	C.cgo_callfn_GDExtensionSpxAudioSetAudioVolume(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxAudioGetAudioVolume(
	) GdFloat {
	arg0 := (C.GDExtensionSpxAudioGetAudioVolume)(FFI.SpxAudioGetAudioVolume)
	ret := C.cgo_callfn_GDExtensionSpxAudioGetAudioVolume(arg0,)
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxAudioIsMusicPlaying(
	) GdBool {
	arg0 := (C.GDExtensionSpxAudioIsMusicPlaying)(FFI.SpxAudioIsMusicPlaying)
	ret := C.cgo_callfn_GDExtensionSpxAudioIsMusicPlaying(arg0,)
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxAudioPlayMusic(
	path GdString,
	)  {
	arg0 := (C.GDExtensionSpxAudioPlayMusic)(FFI.SpxAudioPlayMusic)
	arg1 := (C.GdString)(path)
	C.cgo_callfn_GDExtensionSpxAudioPlayMusic(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxAudioSetMusicVolume(
	volume GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxAudioSetMusicVolume)(FFI.SpxAudioSetMusicVolume)
	arg1 := (C.GdFloat)(volume)
	C.cgo_callfn_GDExtensionSpxAudioSetMusicVolume(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxAudioGetMusicVolume(
	) GdFloat {
	arg0 := (C.GDExtensionSpxAudioGetMusicVolume)(FFI.SpxAudioGetMusicVolume)
	ret := C.cgo_callfn_GDExtensionSpxAudioGetMusicVolume(arg0,)
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxAudioPauseMusic(
	)  {
	arg0 := (C.GDExtensionSpxAudioPauseMusic)(FFI.SpxAudioPauseMusic)
	C.cgo_callfn_GDExtensionSpxAudioPauseMusic(arg0,)
}
func CallFunc_GDExtensionSpxAudioResumeMusic(
	)  {
	arg0 := (C.GDExtensionSpxAudioResumeMusic)(FFI.SpxAudioResumeMusic)
	C.cgo_callfn_GDExtensionSpxAudioResumeMusic(arg0,)
}
func CallFunc_GDExtensionSpxAudioGetMusicTimer(
	) GdFloat {
	arg0 := (C.GDExtensionSpxAudioGetMusicTimer)(FFI.SpxAudioGetMusicTimer)
	ret := C.cgo_callfn_GDExtensionSpxAudioGetMusicTimer(arg0,)
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxAudioSetMusicTimer(
	time GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxAudioSetMusicTimer)(FFI.SpxAudioSetMusicTimer)
	arg1 := (C.GdFloat)(time)
	C.cgo_callfn_GDExtensionSpxAudioSetMusicTimer(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxPhysicSetGravity(
	gravity GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetGravity)(FFI.SpxPhysicSetGravity)
	arg1 := (C.GdFloat)(gravity)
	C.cgo_callfn_GDExtensionSpxPhysicSetGravity(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxPhysicGetGravity(
	) GdFloat {
	arg0 := (C.GDExtensionSpxPhysicGetGravity)(FFI.SpxPhysicGetGravity)
	ret := C.cgo_callfn_GDExtensionSpxPhysicGetGravity(arg0,)
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxPhysicSetVelocity(
	id GdInt,
	velocity GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetVelocity)(FFI.SpxPhysicSetVelocity)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(velocity)
	C.cgo_callfn_GDExtensionSpxPhysicSetVelocity(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicGetVelocity(
	id GdInt,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxPhysicGetVelocity)(FFI.SpxPhysicGetVelocity)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicGetVelocity(arg0,arg1,)
	
	return (GdVec2)(ret)
}
func CallFunc_GDExtensionSpxPhysicSetMass(
	id GdInt,
	mass GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetMass)(FFI.SpxPhysicSetMass)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdFloat)(mass)
	C.cgo_callfn_GDExtensionSpxPhysicSetMass(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicGetMass(
	id GdInt,
	) GdFloat {
	arg0 := (C.GDExtensionSpxPhysicGetMass)(FFI.SpxPhysicGetMass)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicGetMass(arg0,arg1,)
	
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxPhysicAddForce(
	id GdInt,
	force GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxPhysicAddForce)(FFI.SpxPhysicAddForce)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(force)
	C.cgo_callfn_GDExtensionSpxPhysicAddForce(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicAddImpulse(
	id GdInt,
	impulse GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxPhysicAddImpulse)(FFI.SpxPhysicAddImpulse)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(impulse)
	C.cgo_callfn_GDExtensionSpxPhysicAddImpulse(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicSetCollisionLayer(
	id GdInt,
	layer GdInt,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetCollisionLayer)(FFI.SpxPhysicSetCollisionLayer)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdInt)(layer)
	C.cgo_callfn_GDExtensionSpxPhysicSetCollisionLayer(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicGetCollisionLayer(
	id GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxPhysicGetCollisionLayer)(FFI.SpxPhysicGetCollisionLayer)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicGetCollisionLayer(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxPhysicSetCollisionMask(
	id GdInt,
	mask GdInt,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetCollisionMask)(FFI.SpxPhysicSetCollisionMask)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdInt)(mask)
	C.cgo_callfn_GDExtensionSpxPhysicSetCollisionMask(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicGetCollisionMask(
	id GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxPhysicGetCollisionMask)(FFI.SpxPhysicGetCollisionMask)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicGetCollisionMask(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxPhysicGetColliderType(
	id GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxPhysicGetColliderType)(FFI.SpxPhysicGetColliderType)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicGetColliderType(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxPhysicAddColliderRect(
	id GdInt,
	center GdVec2,
	size GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxPhysicAddColliderRect)(FFI.SpxPhysicAddColliderRect)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(center)
	arg3 := (C.GdVec2)(size)
	C.cgo_callfn_GDExtensionSpxPhysicAddColliderRect(arg0,arg1,arg2,arg3,)
	
	
	
}
func CallFunc_GDExtensionSpxPhysicAddColliderCircle(
	id GdInt,
	center GdVec2,
	radius GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxPhysicAddColliderCircle)(FFI.SpxPhysicAddColliderCircle)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(center)
	arg3 := (C.GdFloat)(radius)
	C.cgo_callfn_GDExtensionSpxPhysicAddColliderCircle(arg0,arg1,arg2,arg3,)
	
	
	
}
func CallFunc_GDExtensionSpxPhysicAddColliderCapsule(
	id GdInt,
	center GdVec2,
	size GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxPhysicAddColliderCapsule)(FFI.SpxPhysicAddColliderCapsule)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(center)
	arg3 := (C.GdVec2)(size)
	C.cgo_callfn_GDExtensionSpxPhysicAddColliderCapsule(arg0,arg1,arg2,arg3,)
	
	
	
}
func CallFunc_GDExtensionSpxPhysicSetTrigger(
	id GdInt,
	trigger GdBool,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetTrigger)(FFI.SpxPhysicSetTrigger)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdBool)(trigger)
	C.cgo_callfn_GDExtensionSpxPhysicSetTrigger(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicIsTrigger(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxPhysicIsTrigger)(FFI.SpxPhysicIsTrigger)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicIsTrigger(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxPhysicSetCollisionEnabled(
	id GdInt,
	enabled GdBool,
	)  {
	arg0 := (C.GDExtensionSpxPhysicSetCollisionEnabled)(FFI.SpxPhysicSetCollisionEnabled)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdBool)(enabled)
	C.cgo_callfn_GDExtensionSpxPhysicSetCollisionEnabled(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxPhysicIsCollisionEnabled(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxPhysicIsCollisionEnabled)(FFI.SpxPhysicIsCollisionEnabled)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxPhysicIsCollisionEnabled(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxSpriteCreateSprite(
	path GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteCreateSprite)(FFI.SpxSpriteCreateSprite)
	arg1 := (C.GdString)(path)
	ret := C.cgo_callfn_GDExtensionSpxSpriteCreateSprite(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxSpriteCloneSprite(
	id GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteCloneSprite)(FFI.SpxSpriteCloneSprite)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteCloneSprite(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxSpriteDestroySprite(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteDestroySprite)(FFI.SpxSpriteDestroySprite)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteDestroySprite(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxSpriteIsSpriteAlive(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsSpriteAlive)(FFI.SpxSpriteIsSpriteAlive)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteIsSpriteAlive(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxSpriteSetPosition(
	id GdInt,
	pos GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetPosition)(FFI.SpxSpriteSetPosition)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(pos)
	C.cgo_callfn_GDExtensionSpxSpriteSetPosition(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxSpriteSetRotation(
	id GdInt,
	rot GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetRotation)(FFI.SpxSpriteSetRotation)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(rot)
	C.cgo_callfn_GDExtensionSpxSpriteSetRotation(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxSpriteSetScale(
	id GdInt,
	scale GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetScale)(FFI.SpxSpriteSetScale)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdVec2)(scale)
	C.cgo_callfn_GDExtensionSpxSpriteSetScale(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxSpriteGetPosition(
	id GdInt,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxSpriteGetPosition)(FFI.SpxSpriteGetPosition)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteGetPosition(arg0,arg1,)
	
	return (GdVec2)(ret)
}
func CallFunc_GDExtensionSpxSpriteGetRotation(
	id GdInt,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxSpriteGetRotation)(FFI.SpxSpriteGetRotation)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteGetRotation(arg0,arg1,)
	
	return (GdVec2)(ret)
}
func CallFunc_GDExtensionSpxSpriteGetScale(
	id GdInt,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxSpriteGetScale)(FFI.SpxSpriteGetScale)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteGetScale(arg0,arg1,)
	
	return (GdVec2)(ret)
}
func CallFunc_GDExtensionSpxSpriteSetColor(
	id GdInt,
	GdColor GdColor,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetColor)(FFI.SpxSpriteSetColor)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdColor)(GdColor)
	C.cgo_callfn_GDExtensionSpxSpriteSetColor(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxSpriteGetColor(
	id GdInt,
	) GdColor {
	arg0 := (C.GDExtensionSpxSpriteGetColor)(FFI.SpxSpriteGetColor)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteGetColor(arg0,arg1,)
	
	return (GdColor)(ret)
}
func CallFunc_GDExtensionSpxSpriteUpdateTexture(
	id GdInt,
	path GdString,
	)  {
	arg0 := (C.GDExtensionSpxSpriteUpdateTexture)(FFI.SpxSpriteUpdateTexture)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdString)(path)
	C.cgo_callfn_GDExtensionSpxSpriteUpdateTexture(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxSpriteGetTexture(
	id GdInt,
	) GdString {
	arg0 := (C.GDExtensionSpxSpriteGetTexture)(FFI.SpxSpriteGetTexture)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteGetTexture(arg0,arg1,)
	
	return (GdString)(ret)
}
func CallFunc_GDExtensionSpxSpriteSetVisible(
	id GdInt,
	visible GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetVisible)(FFI.SpxSpriteSetVisible)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdBool)(visible)
	C.cgo_callfn_GDExtensionSpxSpriteSetVisible(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxSpriteGetVisible(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteGetVisible)(FFI.SpxSpriteGetVisible)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxSpriteGetVisible(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxSpriteUpdateZIndex(
	id GdInt,
	z GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteUpdateZIndex)(FFI.SpxSpriteUpdateZIndex)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdInt)(z)
	C.cgo_callfn_GDExtensionSpxSpriteUpdateZIndex(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUICreateButton(
	path GdString,
	rect GdRect,
	text GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateButton)(FFI.SpxUICreateButton)
	arg1 := (C.GdString)(path)
	arg2 := (C.GdRect)(rect)
	arg3 := (C.GdString)(text)
	ret := C.cgo_callfn_GDExtensionSpxUICreateButton(arg0,arg1,arg2,arg3,)
	
	
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUICreateLabel(
	path GdString,
	rect GdRect,
	text GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateLabel)(FFI.SpxUICreateLabel)
	arg1 := (C.GdString)(path)
	arg2 := (C.GdRect)(rect)
	arg3 := (C.GdString)(text)
	ret := C.cgo_callfn_GDExtensionSpxUICreateLabel(arg0,arg1,arg2,arg3,)
	
	
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUICreateImage(
	path GdString,
	rect GdRect,
	GdColor GdColor,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateImage)(FFI.SpxUICreateImage)
	arg1 := (C.GdString)(path)
	arg2 := (C.GdRect)(rect)
	arg3 := (C.GdColor)(GdColor)
	ret := C.cgo_callfn_GDExtensionSpxUICreateImage(arg0,arg1,arg2,arg3,)
	
	
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUICreateSlider(
	path GdString,
	rect GdRect,
	value GdFloat,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateSlider)(FFI.SpxUICreateSlider)
	arg1 := (C.GdString)(path)
	arg2 := (C.GdRect)(rect)
	arg3 := (C.GdFloat)(value)
	ret := C.cgo_callfn_GDExtensionSpxUICreateSlider(arg0,arg1,arg2,arg3,)
	
	
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUICreateToggle(
	path GdString,
	rect GdRect,
	value GdBool,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateToggle)(FFI.SpxUICreateToggle)
	arg1 := (C.GdString)(path)
	arg2 := (C.GdRect)(rect)
	arg3 := (C.GdBool)(value)
	ret := C.cgo_callfn_GDExtensionSpxUICreateToggle(arg0,arg1,arg2,arg3,)
	
	
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUICreateInput(
	path GdString,
	rect GdRect,
	text GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateInput)(FFI.SpxUICreateInput)
	arg1 := (C.GdString)(path)
	arg2 := (C.GdRect)(rect)
	arg3 := (C.GdString)(text)
	ret := C.cgo_callfn_GDExtensionSpxUICreateInput(arg0,arg1,arg2,arg3,)
	
	
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUIGetType(
	id GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxUIGetType)(FFI.SpxUIGetType)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetType(arg0,arg1,)
	
	return (GdInt)(ret)
}
func CallFunc_GDExtensionSpxUISetInteractable(
	id GdInt,
	interactable GdBool,
	)  {
	arg0 := (C.GDExtensionSpxUISetInteractable)(FFI.SpxUISetInteractable)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdBool)(interactable)
	C.cgo_callfn_GDExtensionSpxUISetInteractable(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUIGetInteractable(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxUIGetInteractable)(FFI.SpxUIGetInteractable)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetInteractable(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxUISetText(
	id GdInt,
	text GdString,
	)  {
	arg0 := (C.GDExtensionSpxUISetText)(FFI.SpxUISetText)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdString)(text)
	C.cgo_callfn_GDExtensionSpxUISetText(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUIGetText(
	id GdInt,
	) GdString {
	arg0 := (C.GDExtensionSpxUIGetText)(FFI.SpxUIGetText)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetText(arg0,arg1,)
	
	return (GdString)(ret)
}
func CallFunc_GDExtensionSpxUISetRect(
	id GdInt,
	rect GdRect,
	)  {
	arg0 := (C.GDExtensionSpxUISetRect)(FFI.SpxUISetRect)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdRect)(rect)
	C.cgo_callfn_GDExtensionSpxUISetRect(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUIGetRect(
	id GdInt,
	) GdRect {
	arg0 := (C.GDExtensionSpxUIGetRect)(FFI.SpxUIGetRect)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetRect(arg0,arg1,)
	
	return (GdRect)(ret)
}
func CallFunc_GDExtensionSpxUISetColor(
	id GdInt,
	GdColor GdColor,
	)  {
	arg0 := (C.GDExtensionSpxUISetColor)(FFI.SpxUISetColor)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdColor)(GdColor)
	C.cgo_callfn_GDExtensionSpxUISetColor(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUIGetColor(
	id GdInt,
	) GdColor {
	arg0 := (C.GDExtensionSpxUIGetColor)(FFI.SpxUIGetColor)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetColor(arg0,arg1,)
	
	return (GdColor)(ret)
}
func CallFunc_GDExtensionSpxUISetFontSize(
	id GdInt,
	size GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxUISetFontSize)(FFI.SpxUISetFontSize)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdFloat)(size)
	C.cgo_callfn_GDExtensionSpxUISetFontSize(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUIGetFontSize(
	id GdInt,
	) GdFloat {
	arg0 := (C.GDExtensionSpxUIGetFontSize)(FFI.SpxUIGetFontSize)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetFontSize(arg0,arg1,)
	
	return (GdFloat)(ret)
}
func CallFunc_GDExtensionSpxUISetVisible(
	id GdInt,
	visible GdBool,
	)  {
	arg0 := (C.GDExtensionSpxUISetVisible)(FFI.SpxUISetVisible)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdBool)(visible)
	C.cgo_callfn_GDExtensionSpxUISetVisible(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxUIGetVisible(
	id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxUIGetVisible)(FFI.SpxUIGetVisible)
	arg1 := (C.GdInt)(id)
	ret := C.cgo_callfn_GDExtensionSpxUIGetVisible(arg0,arg1,)
	
	return (GdBool)(ret)
}
func CallFunc_GDExtensionSpxCallbackOnEngineStart(
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnEngineStart)(FFI.SpxCallbackOnEngineStart)
	C.cgo_callfn_GDExtensionSpxCallbackOnEngineStart(arg0,)
}
func CallFunc_GDExtensionSpxCallbackOnEngineUpdate(
	delta GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnEngineUpdate)(FFI.SpxCallbackOnEngineUpdate)
	arg1 := (C.GdFloat)(delta)
	C.cgo_callfn_GDExtensionSpxCallbackOnEngineUpdate(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnEngineDestroy(
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnEngineDestroy)(FFI.SpxCallbackOnEngineDestroy)
	C.cgo_callfn_GDExtensionSpxCallbackOnEngineDestroy(arg0,)
}
func CallFunc_GDExtensionSpxCallbackOnSpriteReady(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnSpriteReady)(FFI.SpxCallbackOnSpriteReady)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnSpriteReady(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnSpriteUpdated(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnSpriteUpdated)(FFI.SpxCallbackOnSpriteUpdated)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnSpriteUpdated(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnSpriteDestroyed(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnSpriteDestroyed)(FFI.SpxCallbackOnSpriteDestroyed)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnSpriteDestroyed(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnMousePressed(
	keyid GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnMousePressed)(FFI.SpxCallbackOnMousePressed)
	arg1 := (C.GdInt)(keyid)
	C.cgo_callfn_GDExtensionSpxCallbackOnMousePressed(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnMouseReleased(
	keyid GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnMouseReleased)(FFI.SpxCallbackOnMouseReleased)
	arg1 := (C.GdInt)(keyid)
	C.cgo_callfn_GDExtensionSpxCallbackOnMouseReleased(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnKeyPressed(
	keyid GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnKeyPressed)(FFI.SpxCallbackOnKeyPressed)
	arg1 := (C.GdInt)(keyid)
	C.cgo_callfn_GDExtensionSpxCallbackOnKeyPressed(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnKeyReleased(
	keyid GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnKeyReleased)(FFI.SpxCallbackOnKeyReleased)
	arg1 := (C.GdInt)(keyid)
	C.cgo_callfn_GDExtensionSpxCallbackOnKeyReleased(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnActionPressed(
	action_name GdString,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnActionPressed)(FFI.SpxCallbackOnActionPressed)
	arg1 := (C.GdString)(action_name)
	C.cgo_callfn_GDExtensionSpxCallbackOnActionPressed(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnActionJustPressed(
	action_name GdString,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnActionJustPressed)(FFI.SpxCallbackOnActionJustPressed)
	arg1 := (C.GdString)(action_name)
	C.cgo_callfn_GDExtensionSpxCallbackOnActionJustPressed(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnActionJustReleased(
	action_name GdString,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnActionJustReleased)(FFI.SpxCallbackOnActionJustReleased)
	arg1 := (C.GdString)(action_name)
	C.cgo_callfn_GDExtensionSpxCallbackOnActionJustReleased(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnAxisChanged(
	action_name GdString,
	value GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnAxisChanged)(FFI.SpxCallbackOnAxisChanged)
	arg1 := (C.GdString)(action_name)
	arg2 := (C.GdFloat)(value)
	C.cgo_callfn_GDExtensionSpxCallbackOnAxisChanged(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnCollisionEnter(
	self_id GdInt,
	other_id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnCollisionEnter)(FFI.SpxCallbackOnCollisionEnter)
	arg1 := (C.GdInt)(self_id)
	arg2 := (C.GdInt)(other_id)
	C.cgo_callfn_GDExtensionSpxCallbackOnCollisionEnter(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnCollisionStay(
	self_id GdInt,
	other_id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnCollisionStay)(FFI.SpxCallbackOnCollisionStay)
	arg1 := (C.GdInt)(self_id)
	arg2 := (C.GdInt)(other_id)
	C.cgo_callfn_GDExtensionSpxCallbackOnCollisionStay(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnCollisionExit(
	self_id GdInt,
	other_id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnCollisionExit)(FFI.SpxCallbackOnCollisionExit)
	arg1 := (C.GdInt)(self_id)
	arg2 := (C.GdInt)(other_id)
	C.cgo_callfn_GDExtensionSpxCallbackOnCollisionExit(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnTriggerEnter(
	self_id GdInt,
	other_id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnTriggerEnter)(FFI.SpxCallbackOnTriggerEnter)
	arg1 := (C.GdInt)(self_id)
	arg2 := (C.GdInt)(other_id)
	C.cgo_callfn_GDExtensionSpxCallbackOnTriggerEnter(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnTriggerStay(
	self_id GdInt,
	other_id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnTriggerStay)(FFI.SpxCallbackOnTriggerStay)
	arg1 := (C.GdInt)(self_id)
	arg2 := (C.GdInt)(other_id)
	C.cgo_callfn_GDExtensionSpxCallbackOnTriggerStay(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnTriggerExit(
	self_id GdInt,
	other_id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnTriggerExit)(FFI.SpxCallbackOnTriggerExit)
	arg1 := (C.GdInt)(self_id)
	arg2 := (C.GdInt)(other_id)
	C.cgo_callfn_GDExtensionSpxCallbackOnTriggerExit(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnUIPressed(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnUIPressed)(FFI.SpxCallbackOnUIPressed)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnUIPressed(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnUIReleased(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnUIReleased)(FFI.SpxCallbackOnUIReleased)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnUIReleased(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnUIHovered(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnUIHovered)(FFI.SpxCallbackOnUIHovered)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnUIHovered(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnUIClicked(
	id GdInt,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnUIClicked)(FFI.SpxCallbackOnUIClicked)
	arg1 := (C.GdInt)(id)
	C.cgo_callfn_GDExtensionSpxCallbackOnUIClicked(arg0,arg1,)
	
}
func CallFunc_GDExtensionSpxCallbackOnUIToggle(
	id GdInt,
	is_on GdBool,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnUIToggle)(FFI.SpxCallbackOnUIToggle)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdBool)(is_on)
	C.cgo_callfn_GDExtensionSpxCallbackOnUIToggle(arg0,arg1,arg2,)
	
	
}
func CallFunc_GDExtensionSpxCallbackOnUITextChanged(
	id GdInt,
	text GdString,
	)  {
	arg0 := (C.GDExtensionSpxCallbackOnUITextChanged)(FFI.SpxCallbackOnUITextChanged)
	arg1 := (C.GdInt)(id)
	arg2 := (C.GdString)(text)
	C.cgo_callfn_GDExtensionSpxCallbackOnUITextChanged(arg0,arg1,arg2,)
	
	
}