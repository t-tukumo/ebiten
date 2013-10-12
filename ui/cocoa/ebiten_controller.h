// -*- objc -*-

#ifndef GO_EBITEN_UI_COCOA_EBITEN_CONTROLLER_H_
#define GO_EBITEN_UI_COCOA_EBITEN_CONTROLLER_H_

#include <Cocoa/Cocoa.h>

@interface EbitenController : NSObject<NSApplicationDelegate>

- (id)initWithWindow:(NSWindow*)window;
- (void)applicationDidFinishLaunching:(NSNotification*)aNotification;

@end

#endif
