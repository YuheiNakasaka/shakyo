#!/bin/bash
#

input="cat_sozai.gif"
output="/tmp/anim20151129.gif"
pingpong=0                          # comment this if you dont want a pingpong loop
frame_start=1                       # which frame to pick as starting frame from inputfile?
frame_stop=57                       # and which frame to stop?
frame_bg=1                          # which frame should function as background 'still' ?
overlay_x=320                       # which part should be animated?
overlay_y=150                        #
overlay_w=60                       #
overlay_h=65                       #
viewport_x=0                        # crop final output
viewport_y=0                        #
viewport_w=411                      #
viewport_h=315                      #
speed_current=0
declare -A speed

# framespeed script
speed[0]=4
speed[1]=4
speed[6]=4
speed[7]=4

init(){
  [[ ! -d /tmp/frames ]] && mkdir frames
  rm frames/* &>/dev/null
}

movie_to_frames(){
  echo "extracting frame $frame_bg as background frame"
  convert "$input[$frame_bg]" /tmp/frames/bg.png
  convert "$input" -repage 0x0 -crop $overlay_w"x"$overlay_h+$overlay_x+$overlay_y +repage /tmp/frames/frame.png
  echo "+/- $(ls /tmp/frames/* | wc -l) frames generated"
  #echo "cropping the animationpart"
  #for file in /tmp/frames/frame-*.png; do
  #  mogrify -crop 8x8 "$file"
  #done
}

iterate_frames_reversed(){
  start="$1"; stop="$2"; offset="$3"
  for((i=start-1;i>stop;i--)); do
    [[ ${#speed[$offset]} > 0 ]] && speed_current=${speed[$offset]}
    echo "processing reversed frame $offset (source:$i) with speed $speed_current"
    frames="$frames -delay $speed_current /tmp/frames/frame-out-$offset.png"
    convert "/tmp/frames/bg.png" /tmp/frames/frame-$i.png -geometry +$overlay_x+$overlay_y \
      -compose over -composite /tmp/frames/frame-out-$offset.png
    ((offset+=1))
  done
}

iterate_frames(){
  start="$1"; stop="$2"; offset=0
  for((i=start;i<stop;i++)); do
    [[ ${#speed[$offset]} > 0 ]] && speed_current=${speed[$offset]}
    echo "processing frame $offset (source:$i) with speed $speed_current"
    frames="$frames -delay $speed_current /tmp/frames/frame-out-$offset.png"
    convert "/tmp/frames/bg.png" /tmp/frames/frame-$i.png -geometry +$overlay_x+$overlay_y\
      -compose over -composite /tmp/frames/frame-out-$offset.png
    ((offset+=1))
  done
}

frames_to_gif(){
  frames=""
  iterate_frames $frame_start $frame_stop
  echo ${frames}
  # [[ -n $pingpong ]] && iterate_frames_reversed $frame_stop $frame_start $((frame_stop-frame_start))
  convert ${frames} -loop 0 -crop $viewport_w"x"$viewport_h+$viewport_x+$viewport_y -layers Optimize "$output"
}

filter(){
  [[ -n DEBUG ]] && cat - || cat - | while read line; do tput el; printf "\r$line"; done
}

init ; movie_to_frames; frames_to_gif | filter
echo ""
ls -la "$output"