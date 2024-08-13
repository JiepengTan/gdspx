﻿/**************************************************************************/
/*  gdextension_spx_ext.h                                                 */
/**************************************************************************/
/*                         This file is part of:                          */
/*                             GODOT ENGINE                               */
/*                        https://godotengine.org                         */
/**************************************************************************/
/* Copyright (c) 2014-present Godot Engine contributors (see AUTHORS.md). */
/* Copyright (c) 2007-2014 Juan Linietsky, Ariel Manzur.                  */
/*                                                                        */
/* Permission is hereby granted, free of charge, to any person obtaining  */
/* a copy of this software and associated documentation files (the        */
/* "Software"), to deal in the Software without restriction, including    */
/* without limitation the rights to use, copy, modify, merge, publish,    */
/* distribute, sublicense, and/or sell copies of the Software, and to     */
/* permit persons to whom the Software is furnished to do so, subject to  */
/* the following conditions:                                              */
/*                                                                        */
/* The above copyright notice and this permission notice shall be         */
/* included in all copies or substantial portions of the Software.        */
/*                                                                        */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,        */
/* EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF     */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. */
/* IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,   */
/* TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE      */
/* SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                 */
/**************************************************************************/
#ifndef GDEXTENSION_SPX_EXT_H
#define GDEXTENSION_SPX_EXT_H

#include "gdextension_interface.h"
#ifndef NOT_GODOT_ENGINE
#include "core/variant/variant.h"
extern void gdextension_spx_setup_interface();
#endif

#ifdef __cplusplus
extern "C" {
#endif

typedef GDExtensionConstStringPtr GdString;
typedef GDExtensionInt GdInt;
typedef GDExtensionInt GdObj;
typedef GDExtensionBool GdBool;
typedef real_t GdFloat;
typedef Vector4 GdVec4;
typedef Vector3 GdVec3;
typedef Vector2 GdVec2;
typedef Color GdColor;
typedef Rect2 GdRect2;

typedef void *GDExtensionSpxCallbackInfoPtr;
typedef void (*GDExtensionSpxGlobalRegisterCallbacks)(GDExtensionSpxCallbackInfoPtr callback_ptr);

typedef GdVec2 (*GDExtensionSpxInputGetMousePos)();
typedef GdBool (*GDExtensionSpxInputGetMouseState)(GdInt obj);
typedef GdInt (*GDExtensionSpxInputGetKeyState)(GdInt key);
typedef GdFloat (*GDExtensionSpxInputGetAxis)(GdString axis);
typedef GdBool (*GDExtensionSpxInputIsActionPressed)(GdString action);
typedef GdBool (*GDExtensionSpxInputIsActionJustPressed)(GdString action);
typedef GdBool (*GDExtensionSpxInputIsActionJustReleased)(GdString action);

// audio manager
typedef void (*GDExtensionSpxAudioPlayAudio)(GdString path);
typedef void (*GDExtensionSpxAudioSetAudioVolume)(GdFloat volume);
typedef GdFloat (*GDExtensionSpxAudioGetAudioVolume)();

typedef GdBool (*GDExtensionSpxAudioIsMusicPlaying)();
typedef void (*GDExtensionSpxAudioPlayMusic)(GdString path);
typedef void (*GDExtensionSpxAudioSetMusicVolume)(GdFloat volume);
typedef GdFloat (*GDExtensionSpxAudioGetMusicVolume)();
typedef void (*GDExtensionSpxAudioPauseMusic)();
typedef void (*GDExtensionSpxAudioResumeMusic)();
typedef GdFloat (*GDExtensionSpxAudioGetMusicTimer)();
typedef void (*GDExtensionSpxAudioSetMusicTimer)(GdFloat time);

// physic
typedef void (*GDExtensionSpxPhysicSetGravity)(GdFloat gravity);
typedef GdFloat (*GDExtensionSpxPhysicGetGravity)();

typedef void (*GDExtensionSpxPhysicSetVelocity)(GdObj obj, GdVec2 velocity);
typedef GdVec2 (*GDExtensionSpxPhysicGetVelocity)(GdObj obj);

typedef void (*GDExtensionSpxPhysicSetMass)(GdObj obj, GdFloat mass);
typedef GdFloat (*GDExtensionSpxPhysicGetMass)(GdObj obj);

typedef void (*GDExtensionSpxPhysicAddForce)(GdObj obj, GdVec2 force);
typedef void (*GDExtensionSpxPhysicAddImpulse)(GdObj obj, GdVec2 impulse);
typedef void (*GDExtensionSpxPhysicSetCollisionLayer)(GdObj obj, GdInt layer);

typedef GdInt (*GDExtensionSpxPhysicGetCollisionLayer)(GdObj obj);
typedef void (*GDExtensionSpxPhysicSetCollisionMask)(GdObj obj, GdInt mask);
typedef GdInt (*GDExtensionSpxPhysicGetCollisionMask)(GdObj obj);

typedef GdInt (*GDExtensionSpxPhysicGetColliderType)(GdObj obj);
typedef void (*GDExtensionSpxPhysicAddColliderRect)(GdObj obj, GdVec2 center, GdVec2 size);
typedef void (*GDExtensionSpxPhysicAddColliderCircle)(GdObj obj, GdVec2 center, GdFloat radius);
typedef void (*GDExtensionSpxPhysicAddColliderCapsule)(GdObj obj, GdVec2 center, GdVec2 size);

typedef void (*GDExtensionSpxPhysicSetTrigger)(GdObj obj, GdBool trigger);
typedef GdBool (*GDExtensionSpxPhysicIsTrigger)(GdObj obj);

typedef void (*GDExtensionSpxPhysicSetCollisionEnabled)(GdObj obj, GdBool enabled);
typedef GdBool (*GDExtensionSpxPhysicIsCollisionEnabled)(GdObj obj);

// sprite
typedef GdObj (*GDExtensionSpxSpriteCreateSprite)(GdString path);
typedef GdObj (*GDExtensionSpxSpriteCloneSprite)(GdObj obj);
typedef GdBool (*GDExtensionSpxSpriteDestroySprite)(GdObj obj);
typedef GdBool (*GDExtensionSpxSpriteIsSpriteAlive)(GdObj obj);

typedef void (*GDExtensionSpxSpriteSetPosition)(GdObj obj, GdVec2 pos);
typedef void (*GDExtensionSpxSpriteSetRotation)(GdObj obj, GdFloat rot);
typedef void (*GDExtensionSpxSpriteSetScale)(GdObj obj, GdVec2 scale);

typedef GdVec2 (*GDExtensionSpxSpriteGetPosition)(GdObj obj);
typedef GdFloat (*GDExtensionSpxSpriteGetRotation)(GdObj obj);
typedef GdVec2 (*GDExtensionSpxSpriteGetScale)(GdObj obj);

typedef void (*GDExtensionSpxSpriteSetColor)(GdObj obj, GdColor color);
typedef GdColor (*GDExtensionSpxSpriteGetColor)(GdObj obj);

typedef void (*GDExtensionSpxSpriteUpdateTexture)(GdObj obj, GdString path);
typedef GdString (*GDExtensionSpxSpriteGetTexture)(GdObj obj);

typedef void (*GDExtensionSpxSpriteSetVisible)(GdObj obj, GdBool visible);
typedef GdBool (*GDExtensionSpxSpriteGetVisible)(GdObj obj);
typedef void (*GDExtensionSpxSpriteUpdateZIndex)(GdObj obj, GdInt z);

// ui
typedef GdObj (*GDExtensionSpxUICreateButton)(GdString path, GdRect2 rect, GdString text);
typedef GdObj (*GDExtensionSpxUICreateLabel)(GdString path, GdRect2 rect, GdString text);
typedef GdObj (*GDExtensionSpxUICreateImage)(GdString path, GdRect2 rect, GdColor color);
typedef GdObj (*GDExtensionSpxUICreateSlider)(GdString path, GdRect2 rect, GdFloat value);
typedef GdObj (*GDExtensionSpxUICreateToggle)(GdString path, GdRect2 rect, GdBool value);
typedef GdObj (*GDExtensionSpxUICreateInput)(GdString path, GdRect2 rect, GdString text);

typedef GdInt (*GDExtensionSpxUIGetType)(GdObj obj);
typedef void (*GDExtensionSpxUISetInteractable)(GdObj obj, GdBool interactable);
typedef GdBool (*GDExtensionSpxUIGetInteractable)(GdObj obj);

typedef void (*GDExtensionSpxUISetText)(GdObj obj, GdString text);
typedef GdString (*GDExtensionSpxUIGetText)(GdObj obj);

typedef void (*GDExtensionSpxUISetRect)(GdObj obj, GdRect2 rect);
typedef GdRect2 (*GDExtensionSpxUIGetRect)(GdObj obj);

typedef void (*GDExtensionSpxUISetColor)(GdObj obj, GdColor color);
typedef GdColor (*GDExtensionSpxUIGetColor)(GdObj obj);

typedef void (*GDExtensionSpxUISetFontSize)(GdObj obj, GdFloat size);
typedef GdFloat (*GDExtensionSpxUIGetFontSize)(GdObj obj);

typedef void (*GDExtensionSpxUISetVisible)(GdObj obj, GdBool visible);
typedef GdBool (*GDExtensionSpxUIGetVisible)(GdObj obj);

// callback
typedef void (*GDExtensionSpxCallbackOnEngineStart)();
typedef void (*GDExtensionSpxCallbackOnEngineUpdate)(GdFloat delta);
typedef void (*GDExtensionSpxCallbackOnEngineDestroy)();

typedef void (*GDExtensionSpxCallbackOnSpriteReady)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnSpriteUpdated)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnSpriteDestroyed)(GdObj obj);

typedef void (*GDExtensionSpxCallbackOnMousePressed)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnMouseReleased)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnKeyPressed)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnKeyReleased)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnActionPressed)(GdString action_name);
typedef void (*GDExtensionSpxCallbackOnActionJustPressed)(GdString action_name);
typedef void (*GDExtensionSpxCallbackOnActionJustReleased)(GdString action_name);
typedef void (*GDExtensionSpxCallbackOnAxisChanged)(GdString action_name, GdFloat value);

typedef void (*GDExtensionSpxCallbackOnCollisionEnter)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnCollisionStay)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnCollisionExit)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnTriggerEnter)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnTriggerStay)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnTriggerExit)(GdInt self_id, GdInt other_id);

typedef void (*GDExtensionSpxCallbackOnUIPressed)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIReleased)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIHovered)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIClicked)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIToggle)(GdObj obj, GdBool is_on);
typedef void (*GDExtensionSpxCallbackOnUITextChanged)(GdObj obj, GdString text);

// string
typedef void (*GDExtensionSpxStringNewWithLatin1Chars)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents);
typedef void (*GDExtensionSpxStringNewWithUtf8Chars)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents);
typedef void (*GDExtensionSpxStringNewWithLatin1CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GdInt p_size);
typedef void (*GDExtensionSpxStringNewWithUtf8CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GdInt p_size);
typedef GdInt (*GDExtensionSpxStringToLatin1Chars)(GDExtensionConstStringPtr p_self, char *r_text, GdInt p_max_write_length);
typedef GdInt (*GDExtensionSpxStringToUtf8Chars)(GDExtensionConstStringPtr p_self, char *r_text, GdInt p_max_write_length);
// variant
typedef GDExtensionPtrConstructor (*GDExtensionSpxVariantGetPtrConstructor)(GDExtensionVariantType p_type, int32_t p_constructor);
typedef GDExtensionPtrDestructor (*GDExtensionSpxVariantGetPtrDestructor)(GDExtensionVariantType p_type);

typedef struct {
	// engine
	GDExtensionSpxCallbackOnEngineStart func_on_engine_start;
	GDExtensionSpxCallbackOnEngineUpdate func_on_engine_update;
	GDExtensionSpxCallbackOnEngineDestroy func_on_engine_destroy;

	// sprite
	GDExtensionSpxCallbackOnSpriteReady func_on_sprite_ready;
	GDExtensionSpxCallbackOnSpriteUpdated func_on_sprite_updated;
	GDExtensionSpxCallbackOnSpriteDestroyed func_on_sprite_destroyed;

	// input
	GDExtensionSpxCallbackOnMousePressed func_on_mouse_pressed;
	GDExtensionSpxCallbackOnMouseReleased func_on_mouse_released;
	GDExtensionSpxCallbackOnKeyPressed func_on_key_pressed;
	GDExtensionSpxCallbackOnKeyReleased func_on_key_released;
	GDExtensionSpxCallbackOnActionPressed func_on_action_pressed;
	GDExtensionSpxCallbackOnActionJustPressed func_on_action_just_pressed;
	GDExtensionSpxCallbackOnActionJustReleased func_on_action_just_released;
	GDExtensionSpxCallbackOnAxisChanged func_on_axis_changed;

	// physic
	GDExtensionSpxCallbackOnCollisionEnter func_on_collision_enter;
	GDExtensionSpxCallbackOnCollisionStay func_on_collision_stay;
	GDExtensionSpxCallbackOnCollisionExit func_on_collision_exit;
	GDExtensionSpxCallbackOnTriggerEnter func_on_trigger_enter;
	GDExtensionSpxCallbackOnTriggerStay func_on_trigger_stay;
	GDExtensionSpxCallbackOnTriggerExit func_on_trigger_exit;

	// ui
	GDExtensionSpxCallbackOnUIPressed func_on_ui_pressed;
	GDExtensionSpxCallbackOnUIReleased func_on_ui_released;
	GDExtensionSpxCallbackOnUIHovered func_on_ui_hovered;
	GDExtensionSpxCallbackOnUIClicked func_on_ui_clicked;
	GDExtensionSpxCallbackOnUIToggle func_on_ui_toggle;
	GDExtensionSpxCallbackOnUITextChanged func_on_ui_text_changed;

} SpxCallbackInfo;

#ifdef __cplusplus
}
#endif

#endif // GDEXTENSION_SPX_EXT_H
