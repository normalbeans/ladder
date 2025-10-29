channels: 
statechannel -> global state channel to handle rerendering 
compeventchannel -> sends input to components from stdin 
compeventdonechannel -> stops input to component and returns control back ( because component may take more than one input, so each component should handle the input sent to it and stop when it deems it is done)

globalstate has 2 main parts : mode (component/page), activecomponent
 
main routine -> just sets things up here .each component is initated with id. 
this is inmportant since the event sent in compeventchannel si bundled with the current active component id. also each component starts its own routine which listens on compeventchannel.
 package initiation is handled in main. 
 alls package.render() // not a routine yet 
 
package.render() - starts a go routine to listen for statechannel and rerender on change also starts input listener which is a loop that blocks using os.read on stdin 

input listener: 
    if mode == page, listen normally for inputs ( enter, up arrow, etc. also sets active component) 
    if mode == component , read input byte by byte and bundle it and send to compeventchannel with activecomponent id .  
also has a select case to listen on compeventdonechannel to switch mode to page


since each component is listeninng in a routine, if the id matches, it does its input handling. after it is done ( may be a special input is received or something), it sends signal on compeventdonechannel to give up control