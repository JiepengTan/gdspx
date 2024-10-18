/*
------------------------------------------------------------------------------
//   This code was generated by template ffi.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package webffi

import (
	"syscall/js"
)

//revive:disable

var (
	API GDExtensionInterface
)

type GDExtensionInterface struct {
	// All of the GDExtension interface functions.
	SpxAudioStopAll                     js.Value
	SpxAudioPlaySfx                     js.Value
	SpxAudioPlayMusic                   js.Value
	SpxAudioPauseMusic                  js.Value
	SpxAudioResumeMusic                 js.Value
	SpxAudioGetMusicTimer               js.Value
	SpxAudioSetMusicTimer               js.Value
	SpxAudioIsMusicPlaying              js.Value
	SpxAudioSetSfxVolume                js.Value
	SpxAudioGetSfxVolume                js.Value
	SpxAudioSetMusicVolume              js.Value
	SpxAudioGetMusicVolume              js.Value
	SpxAudioSetMasterVolume             js.Value
	SpxAudioGetMasterVolume             js.Value
	SpxCameraGetCameraPosition          js.Value
	SpxCameraSetCameraPosition          js.Value
	SpxCameraGetCameraZoom              js.Value
	SpxCameraSetCameraZoom              js.Value
	SpxCameraGetViewportRect            js.Value
	SpxInputGetMousePos                 js.Value
	SpxInputGetKey                      js.Value
	SpxInputGetMouseState               js.Value
	SpxInputGetKeyState                 js.Value
	SpxInputGetAxis                     js.Value
	SpxInputIsActionPressed             js.Value
	SpxInputIsActionJustPressed         js.Value
	SpxInputIsActionJustReleased        js.Value
	SpxPhysicRaycast                    js.Value
	SpxPhysicCheckCollision             js.Value
	SpxPhysicCheckTouchedCameraBoundary js.Value
	SpxPlatformSetWindowSize            js.Value
	SpxPlatformGetWindowSize            js.Value
	SpxPlatformSetWindowTitle           js.Value
	SpxPlatformGetWindowTitle           js.Value
	SpxPlatformSetWindowFullscreen      js.Value
	SpxPlatformIsWindowFullscreen       js.Value
	SpxPlatformSetDebugMode             js.Value
	SpxPlatformIsDebugMode              js.Value
	SpxResGetImageSize                  js.Value
	SpxResReadAllText                   js.Value
	SpxResHasFile                       js.Value
	SpxSceneChangeSceneToFile           js.Value
	SpxSceneReloadCurrentScene          js.Value
	SpxSceneUnloadCurrentScene          js.Value
	SpxSpriteSetDontDestroyOnLoad       js.Value
	SpxSpriteSetProcess                 js.Value
	SpxSpriteSetPhysicProcess           js.Value
	SpxSpriteSetChildPosition           js.Value
	SpxSpriteGetChildPosition           js.Value
	SpxSpriteSetChildRotation           js.Value
	SpxSpriteGetChildRotation           js.Value
	SpxSpriteSetChildScale              js.Value
	SpxSpriteGetChildScale              js.Value
	SpxSpriteCheckCollision             js.Value
	SpxSpriteCheckCollisionWithPoint    js.Value
	SpxSpriteCreateSprite               js.Value
	SpxSpriteCloneSprite                js.Value
	SpxSpriteDestroySprite              js.Value
	SpxSpriteIsSpriteAlive              js.Value
	SpxSpriteSetPosition                js.Value
	SpxSpriteGetPosition                js.Value
	SpxSpriteSetRotation                js.Value
	SpxSpriteGetRotation                js.Value
	SpxSpriteSetScale                   js.Value
	SpxSpriteGetScale                   js.Value
	SpxSpriteSetRenderScale             js.Value
	SpxSpriteGetRenderScale             js.Value
	SpxSpriteSetColor                   js.Value
	SpxSpriteGetColor                   js.Value
	SpxSpriteSetTextureAltas            js.Value
	SpxSpriteSetTexture                 js.Value
	SpxSpriteGetTexture                 js.Value
	SpxSpriteSetVisible                 js.Value
	SpxSpriteGetVisible                 js.Value
	SpxSpriteGetZIndex                  js.Value
	SpxSpriteSetZIndex                  js.Value
	SpxSpritePlayAnim                   js.Value
	SpxSpritePlayBackwardsAnim          js.Value
	SpxSpritePauseAnim                  js.Value
	SpxSpriteStopAnim                   js.Value
	SpxSpriteIsPlayingAnim              js.Value
	SpxSpriteSetAnim                    js.Value
	SpxSpriteGetAnim                    js.Value
	SpxSpriteSetAnimFrame               js.Value
	SpxSpriteGetAnimFrame               js.Value
	SpxSpriteSetAnimSpeedScale          js.Value
	SpxSpriteGetAnimSpeedScale          js.Value
	SpxSpriteGetAnimPlayingSpeed        js.Value
	SpxSpriteSetAnimCentered            js.Value
	SpxSpriteIsAnimCentered             js.Value
	SpxSpriteSetAnimOffset              js.Value
	SpxSpriteGetAnimOffset              js.Value
	SpxSpriteSetAnimFlipH               js.Value
	SpxSpriteIsAnimFlippedH             js.Value
	SpxSpriteSetAnimFlipV               js.Value
	SpxSpriteIsAnimFlippedV             js.Value
	SpxSpriteSetVelocity                js.Value
	SpxSpriteGetVelocity                js.Value
	SpxSpriteIsOnFloor                  js.Value
	SpxSpriteIsOnFloorOnly              js.Value
	SpxSpriteIsOnWall                   js.Value
	SpxSpriteIsOnWallOnly               js.Value
	SpxSpriteIsOnCeiling                js.Value
	SpxSpriteIsOnCeilingOnly            js.Value
	SpxSpriteGetLastMotion              js.Value
	SpxSpriteGetPositionDelta           js.Value
	SpxSpriteGetFloorNormal             js.Value
	SpxSpriteGetWallNormal              js.Value
	SpxSpriteGetRealVelocity            js.Value
	SpxSpriteMoveAndSlide               js.Value
	SpxSpriteSetGravity                 js.Value
	SpxSpriteGetGravity                 js.Value
	SpxSpriteSetMass                    js.Value
	SpxSpriteGetMass                    js.Value
	SpxSpriteAddForce                   js.Value
	SpxSpriteAddImpulse                 js.Value
	SpxSpriteSetCollisionLayer          js.Value
	SpxSpriteGetCollisionLayer          js.Value
	SpxSpriteSetCollisionMask           js.Value
	SpxSpriteGetCollisionMask           js.Value
	SpxSpriteSetTriggerLayer            js.Value
	SpxSpriteGetTriggerLayer            js.Value
	SpxSpriteSetTriggerMask             js.Value
	SpxSpriteGetTriggerMask             js.Value
	SpxSpriteSetColliderRect            js.Value
	SpxSpriteSetColliderCircle          js.Value
	SpxSpriteSetColliderCapsule         js.Value
	SpxSpriteSetCollisionEnabled        js.Value
	SpxSpriteIsCollisionEnabled         js.Value
	SpxSpriteSetTriggerRect             js.Value
	SpxSpriteSetTriggerCircle           js.Value
	SpxSpriteSetTriggerCapsule          js.Value
	SpxSpriteSetTriggerEnabled          js.Value
	SpxSpriteIsTriggerEnabled           js.Value
	SpxUiBindNode                       js.Value
	SpxUiCreateNode                     js.Value
	SpxUiCreateButton                   js.Value
	SpxUiCreateLabel                    js.Value
	SpxUiCreateImage                    js.Value
	SpxUiCreateToggle                   js.Value
	SpxUiCreateSlider                   js.Value
	SpxUiCreateInput                    js.Value
	SpxUiDestroyNode                    js.Value
	SpxUiGetType                        js.Value
	SpxUiSetText                        js.Value
	SpxUiGetText                        js.Value
	SpxUiSetTexture                     js.Value
	SpxUiGetTexture                     js.Value
	SpxUiSetColor                       js.Value
	SpxUiGetColor                       js.Value
	SpxUiSetFontSize                    js.Value
	SpxUiGetFontSize                    js.Value
	SpxUiSetVisible                     js.Value
	SpxUiGetVisible                     js.Value
	SpxUiSetInteractable                js.Value
	SpxUiGetInteractable                js.Value
	SpxUiSetRect                        js.Value
	SpxUiGetRect                        js.Value
	SpxUiGetLayoutDirection             js.Value
	SpxUiSetLayoutDirection             js.Value
	SpxUiGetLayoutMode                  js.Value
	SpxUiSetLayoutMode                  js.Value
	SpxUiGetAnchorsPreset               js.Value
	SpxUiSetAnchorsPreset               js.Value
	SpxUiGetScale                       js.Value
	SpxUiSetScale                       js.Value
	SpxUiGetPosition                    js.Value
	SpxUiSetPosition                    js.Value
	SpxUiGetSize                        js.Value
	SpxUiSetSize                        js.Value
	SpxUiGetGlobalPosition              js.Value
	SpxUiSetGlobalPosition              js.Value
	SpxUiGetRotation                    js.Value
	SpxUiSetRotation                    js.Value
	SpxUiGetFlip                        js.Value
	SpxUiSetFlip                        js.Value
}

func (x *GDExtensionInterface) loadProcAddresses() {
	x.SpxAudioStopAll = dlsymGD("gdspx_audio_stop_all")
	x.SpxAudioPlaySfx = dlsymGD("gdspx_audio_play_sfx")
	x.SpxAudioPlayMusic = dlsymGD("gdspx_audio_play_music")
	x.SpxAudioPauseMusic = dlsymGD("gdspx_audio_pause_music")
	x.SpxAudioResumeMusic = dlsymGD("gdspx_audio_resume_music")
	x.SpxAudioGetMusicTimer = dlsymGD("gdspx_audio_get_music_timer")
	x.SpxAudioSetMusicTimer = dlsymGD("gdspx_audio_set_music_timer")
	x.SpxAudioIsMusicPlaying = dlsymGD("gdspx_audio_is_music_playing")
	x.SpxAudioSetSfxVolume = dlsymGD("gdspx_audio_set_sfx_volume")
	x.SpxAudioGetSfxVolume = dlsymGD("gdspx_audio_get_sfx_volume")
	x.SpxAudioSetMusicVolume = dlsymGD("gdspx_audio_set_music_volume")
	x.SpxAudioGetMusicVolume = dlsymGD("gdspx_audio_get_music_volume")
	x.SpxAudioSetMasterVolume = dlsymGD("gdspx_audio_set_master_volume")
	x.SpxAudioGetMasterVolume = dlsymGD("gdspx_audio_get_master_volume")
	x.SpxCameraGetCameraPosition = dlsymGD("gdspx_camera_get_camera_position")
	x.SpxCameraSetCameraPosition = dlsymGD("gdspx_camera_set_camera_position")
	x.SpxCameraGetCameraZoom = dlsymGD("gdspx_camera_get_camera_zoom")
	x.SpxCameraSetCameraZoom = dlsymGD("gdspx_camera_set_camera_zoom")
	x.SpxCameraGetViewportRect = dlsymGD("gdspx_camera_get_viewport_rect")
	x.SpxInputGetMousePos = dlsymGD("gdspx_input_get_mouse_pos")
	x.SpxInputGetKey = dlsymGD("gdspx_input_get_key")
	x.SpxInputGetMouseState = dlsymGD("gdspx_input_get_mouse_state")
	x.SpxInputGetKeyState = dlsymGD("gdspx_input_get_key_state")
	x.SpxInputGetAxis = dlsymGD("gdspx_input_get_axis")
	x.SpxInputIsActionPressed = dlsymGD("gdspx_input_is_action_pressed")
	x.SpxInputIsActionJustPressed = dlsymGD("gdspx_input_is_action_just_pressed")
	x.SpxInputIsActionJustReleased = dlsymGD("gdspx_input_is_action_just_released")
	x.SpxPhysicRaycast = dlsymGD("gdspx_physic_raycast")
	x.SpxPhysicCheckCollision = dlsymGD("gdspx_physic_check_collision")
	x.SpxPhysicCheckTouchedCameraBoundary = dlsymGD("gdspx_physic_check_touched_camera_boundary")
	x.SpxPlatformSetWindowSize = dlsymGD("gdspx_platform_set_window_size")
	x.SpxPlatformGetWindowSize = dlsymGD("gdspx_platform_get_window_size")
	x.SpxPlatformSetWindowTitle = dlsymGD("gdspx_platform_set_window_title")
	x.SpxPlatformGetWindowTitle = dlsymGD("gdspx_platform_get_window_title")
	x.SpxPlatformSetWindowFullscreen = dlsymGD("gdspx_platform_set_window_fullscreen")
	x.SpxPlatformIsWindowFullscreen = dlsymGD("gdspx_platform_is_window_fullscreen")
	x.SpxPlatformSetDebugMode = dlsymGD("gdspx_platform_set_debug_mode")
	x.SpxPlatformIsDebugMode = dlsymGD("gdspx_platform_is_debug_mode")
	x.SpxResGetImageSize = dlsymGD("gdspx_res_get_image_size")
	x.SpxResReadAllText = dlsymGD("gdspx_res_read_all_text")
	x.SpxResHasFile = dlsymGD("gdspx_res_has_file")
	x.SpxSceneChangeSceneToFile = dlsymGD("gdspx_scene_change_scene_to_file")
	x.SpxSceneReloadCurrentScene = dlsymGD("gdspx_scene_reload_current_scene")
	x.SpxSceneUnloadCurrentScene = dlsymGD("gdspx_scene_unload_current_scene")
	x.SpxSpriteSetDontDestroyOnLoad = dlsymGD("gdspx_sprite_set_dont_destroy_on_load")
	x.SpxSpriteSetProcess = dlsymGD("gdspx_sprite_set_process")
	x.SpxSpriteSetPhysicProcess = dlsymGD("gdspx_sprite_set_physic_process")
	x.SpxSpriteSetChildPosition = dlsymGD("gdspx_sprite_set_child_position")
	x.SpxSpriteGetChildPosition = dlsymGD("gdspx_sprite_get_child_position")
	x.SpxSpriteSetChildRotation = dlsymGD("gdspx_sprite_set_child_rotation")
	x.SpxSpriteGetChildRotation = dlsymGD("gdspx_sprite_get_child_rotation")
	x.SpxSpriteSetChildScale = dlsymGD("gdspx_sprite_set_child_scale")
	x.SpxSpriteGetChildScale = dlsymGD("gdspx_sprite_get_child_scale")
	x.SpxSpriteCheckCollision = dlsymGD("gdspx_sprite_check_collision")
	x.SpxSpriteCheckCollisionWithPoint = dlsymGD("gdspx_sprite_check_collision_with_point")
	x.SpxSpriteCreateSprite = dlsymGD("gdspx_sprite_create_sprite")
	x.SpxSpriteCloneSprite = dlsymGD("gdspx_sprite_clone_sprite")
	x.SpxSpriteDestroySprite = dlsymGD("gdspx_sprite_destroy_sprite")
	x.SpxSpriteIsSpriteAlive = dlsymGD("gdspx_sprite_is_sprite_alive")
	x.SpxSpriteSetPosition = dlsymGD("gdspx_sprite_set_position")
	x.SpxSpriteGetPosition = dlsymGD("gdspx_sprite_get_position")
	x.SpxSpriteSetRotation = dlsymGD("gdspx_sprite_set_rotation")
	x.SpxSpriteGetRotation = dlsymGD("gdspx_sprite_get_rotation")
	x.SpxSpriteSetScale = dlsymGD("gdspx_sprite_set_scale")
	x.SpxSpriteGetScale = dlsymGD("gdspx_sprite_get_scale")
	x.SpxSpriteSetRenderScale = dlsymGD("gdspx_sprite_set_render_scale")
	x.SpxSpriteGetRenderScale = dlsymGD("gdspx_sprite_get_render_scale")
	x.SpxSpriteSetColor = dlsymGD("gdspx_sprite_set_color")
	x.SpxSpriteGetColor = dlsymGD("gdspx_sprite_get_color")
	x.SpxSpriteSetTextureAltas = dlsymGD("gdspx_sprite_set_texture_altas")
	x.SpxSpriteSetTexture = dlsymGD("gdspx_sprite_set_texture")
	x.SpxSpriteGetTexture = dlsymGD("gdspx_sprite_get_texture")
	x.SpxSpriteSetVisible = dlsymGD("gdspx_sprite_set_visible")
	x.SpxSpriteGetVisible = dlsymGD("gdspx_sprite_get_visible")
	x.SpxSpriteGetZIndex = dlsymGD("gdspx_sprite_get_z_index")
	x.SpxSpriteSetZIndex = dlsymGD("gdspx_sprite_set_z_index")
	x.SpxSpritePlayAnim = dlsymGD("gdspx_sprite_play_anim")
	x.SpxSpritePlayBackwardsAnim = dlsymGD("gdspx_sprite_play_backwards_anim")
	x.SpxSpritePauseAnim = dlsymGD("gdspx_sprite_pause_anim")
	x.SpxSpriteStopAnim = dlsymGD("gdspx_sprite_stop_anim")
	x.SpxSpriteIsPlayingAnim = dlsymGD("gdspx_sprite_is_playing_anim")
	x.SpxSpriteSetAnim = dlsymGD("gdspx_sprite_set_anim")
	x.SpxSpriteGetAnim = dlsymGD("gdspx_sprite_get_anim")
	x.SpxSpriteSetAnimFrame = dlsymGD("gdspx_sprite_set_anim_frame")
	x.SpxSpriteGetAnimFrame = dlsymGD("gdspx_sprite_get_anim_frame")
	x.SpxSpriteSetAnimSpeedScale = dlsymGD("gdspx_sprite_set_anim_speed_scale")
	x.SpxSpriteGetAnimSpeedScale = dlsymGD("gdspx_sprite_get_anim_speed_scale")
	x.SpxSpriteGetAnimPlayingSpeed = dlsymGD("gdspx_sprite_get_anim_playing_speed")
	x.SpxSpriteSetAnimCentered = dlsymGD("gdspx_sprite_set_anim_centered")
	x.SpxSpriteIsAnimCentered = dlsymGD("gdspx_sprite_is_anim_centered")
	x.SpxSpriteSetAnimOffset = dlsymGD("gdspx_sprite_set_anim_offset")
	x.SpxSpriteGetAnimOffset = dlsymGD("gdspx_sprite_get_anim_offset")
	x.SpxSpriteSetAnimFlipH = dlsymGD("gdspx_sprite_set_anim_flip_h")
	x.SpxSpriteIsAnimFlippedH = dlsymGD("gdspx_sprite_is_anim_flipped_h")
	x.SpxSpriteSetAnimFlipV = dlsymGD("gdspx_sprite_set_anim_flip_v")
	x.SpxSpriteIsAnimFlippedV = dlsymGD("gdspx_sprite_is_anim_flipped_v")
	x.SpxSpriteSetVelocity = dlsymGD("gdspx_sprite_set_velocity")
	x.SpxSpriteGetVelocity = dlsymGD("gdspx_sprite_get_velocity")
	x.SpxSpriteIsOnFloor = dlsymGD("gdspx_sprite_is_on_floor")
	x.SpxSpriteIsOnFloorOnly = dlsymGD("gdspx_sprite_is_on_floor_only")
	x.SpxSpriteIsOnWall = dlsymGD("gdspx_sprite_is_on_wall")
	x.SpxSpriteIsOnWallOnly = dlsymGD("gdspx_sprite_is_on_wall_only")
	x.SpxSpriteIsOnCeiling = dlsymGD("gdspx_sprite_is_on_ceiling")
	x.SpxSpriteIsOnCeilingOnly = dlsymGD("gdspx_sprite_is_on_ceiling_only")
	x.SpxSpriteGetLastMotion = dlsymGD("gdspx_sprite_get_last_motion")
	x.SpxSpriteGetPositionDelta = dlsymGD("gdspx_sprite_get_position_delta")
	x.SpxSpriteGetFloorNormal = dlsymGD("gdspx_sprite_get_floor_normal")
	x.SpxSpriteGetWallNormal = dlsymGD("gdspx_sprite_get_wall_normal")
	x.SpxSpriteGetRealVelocity = dlsymGD("gdspx_sprite_get_real_velocity")
	x.SpxSpriteMoveAndSlide = dlsymGD("gdspx_sprite_move_and_slide")
	x.SpxSpriteSetGravity = dlsymGD("gdspx_sprite_set_gravity")
	x.SpxSpriteGetGravity = dlsymGD("gdspx_sprite_get_gravity")
	x.SpxSpriteSetMass = dlsymGD("gdspx_sprite_set_mass")
	x.SpxSpriteGetMass = dlsymGD("gdspx_sprite_get_mass")
	x.SpxSpriteAddForce = dlsymGD("gdspx_sprite_add_force")
	x.SpxSpriteAddImpulse = dlsymGD("gdspx_sprite_add_impulse")
	x.SpxSpriteSetCollisionLayer = dlsymGD("gdspx_sprite_set_collision_layer")
	x.SpxSpriteGetCollisionLayer = dlsymGD("gdspx_sprite_get_collision_layer")
	x.SpxSpriteSetCollisionMask = dlsymGD("gdspx_sprite_set_collision_mask")
	x.SpxSpriteGetCollisionMask = dlsymGD("gdspx_sprite_get_collision_mask")
	x.SpxSpriteSetTriggerLayer = dlsymGD("gdspx_sprite_set_trigger_layer")
	x.SpxSpriteGetTriggerLayer = dlsymGD("gdspx_sprite_get_trigger_layer")
	x.SpxSpriteSetTriggerMask = dlsymGD("gdspx_sprite_set_trigger_mask")
	x.SpxSpriteGetTriggerMask = dlsymGD("gdspx_sprite_get_trigger_mask")
	x.SpxSpriteSetColliderRect = dlsymGD("gdspx_sprite_set_collider_rect")
	x.SpxSpriteSetColliderCircle = dlsymGD("gdspx_sprite_set_collider_circle")
	x.SpxSpriteSetColliderCapsule = dlsymGD("gdspx_sprite_set_collider_capsule")
	x.SpxSpriteSetCollisionEnabled = dlsymGD("gdspx_sprite_set_collision_enabled")
	x.SpxSpriteIsCollisionEnabled = dlsymGD("gdspx_sprite_is_collision_enabled")
	x.SpxSpriteSetTriggerRect = dlsymGD("gdspx_sprite_set_trigger_rect")
	x.SpxSpriteSetTriggerCircle = dlsymGD("gdspx_sprite_set_trigger_circle")
	x.SpxSpriteSetTriggerCapsule = dlsymGD("gdspx_sprite_set_trigger_capsule")
	x.SpxSpriteSetTriggerEnabled = dlsymGD("gdspx_sprite_set_trigger_enabled")
	x.SpxSpriteIsTriggerEnabled = dlsymGD("gdspx_sprite_is_trigger_enabled")
	x.SpxUiBindNode = dlsymGD("gdspx_ui_bind_node")
	x.SpxUiCreateNode = dlsymGD("gdspx_ui_create_node")
	x.SpxUiCreateButton = dlsymGD("gdspx_ui_create_button")
	x.SpxUiCreateLabel = dlsymGD("gdspx_ui_create_label")
	x.SpxUiCreateImage = dlsymGD("gdspx_ui_create_image")
	x.SpxUiCreateToggle = dlsymGD("gdspx_ui_create_toggle")
	x.SpxUiCreateSlider = dlsymGD("gdspx_ui_create_slider")
	x.SpxUiCreateInput = dlsymGD("gdspx_ui_create_input")
	x.SpxUiDestroyNode = dlsymGD("gdspx_ui_destroy_node")
	x.SpxUiGetType = dlsymGD("gdspx_ui_get_type")
	x.SpxUiSetText = dlsymGD("gdspx_ui_set_text")
	x.SpxUiGetText = dlsymGD("gdspx_ui_get_text")
	x.SpxUiSetTexture = dlsymGD("gdspx_ui_set_texture")
	x.SpxUiGetTexture = dlsymGD("gdspx_ui_get_texture")
	x.SpxUiSetColor = dlsymGD("gdspx_ui_set_color")
	x.SpxUiGetColor = dlsymGD("gdspx_ui_get_color")
	x.SpxUiSetFontSize = dlsymGD("gdspx_ui_set_font_size")
	x.SpxUiGetFontSize = dlsymGD("gdspx_ui_get_font_size")
	x.SpxUiSetVisible = dlsymGD("gdspx_ui_set_visible")
	x.SpxUiGetVisible = dlsymGD("gdspx_ui_get_visible")
	x.SpxUiSetInteractable = dlsymGD("gdspx_ui_set_interactable")
	x.SpxUiGetInteractable = dlsymGD("gdspx_ui_get_interactable")
	x.SpxUiSetRect = dlsymGD("gdspx_ui_set_rect")
	x.SpxUiGetRect = dlsymGD("gdspx_ui_get_rect")
	x.SpxUiGetLayoutDirection = dlsymGD("gdspx_ui_get_layout_direction")
	x.SpxUiSetLayoutDirection = dlsymGD("gdspx_ui_set_layout_direction")
	x.SpxUiGetLayoutMode = dlsymGD("gdspx_ui_get_layout_mode")
	x.SpxUiSetLayoutMode = dlsymGD("gdspx_ui_set_layout_mode")
	x.SpxUiGetAnchorsPreset = dlsymGD("gdspx_ui_get_anchors_preset")
	x.SpxUiSetAnchorsPreset = dlsymGD("gdspx_ui_set_anchors_preset")
	x.SpxUiGetScale = dlsymGD("gdspx_ui_get_scale")
	x.SpxUiSetScale = dlsymGD("gdspx_ui_set_scale")
	x.SpxUiGetPosition = dlsymGD("gdspx_ui_get_position")
	x.SpxUiSetPosition = dlsymGD("gdspx_ui_set_position")
	x.SpxUiGetSize = dlsymGD("gdspx_ui_get_size")
	x.SpxUiSetSize = dlsymGD("gdspx_ui_set_size")
	x.SpxUiGetGlobalPosition = dlsymGD("gdspx_ui_get_global_position")
	x.SpxUiSetGlobalPosition = dlsymGD("gdspx_ui_set_global_position")
	x.SpxUiGetRotation = dlsymGD("gdspx_ui_get_rotation")
	x.SpxUiSetRotation = dlsymGD("gdspx_ui_set_rotation")
	x.SpxUiGetFlip = dlsymGD("gdspx_ui_get_flip")
	x.SpxUiSetFlip = dlsymGD("gdspx_ui_set_flip")
}
