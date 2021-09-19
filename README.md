# DevOpsAlarm

Set of a physical alarm when a DevOps incident occurs while on call.

## Motivation

On-call DevOps support and engineers need to be notified ASAP when an incident occurs. Traditionally, they receive an email, text, Slack message, etc. However these means of communication are shared with lower urgency messages, and sometimes it is hard to remember to be on red-alert for new messages while on-call.

To solve this, I've made a device using simple off the self parts and free software to support a physical on-call alarm. Combining an industrial signal light stack, a buzzer, a Raspberry Pi (or any device balena supports), and the balena platform, for under $25 you can make a physical alarm that goes off when an incident occurs while on-call.

This project is written in Go to optimize for lower resource devices like the RPi Zero W.

## Features

- Support various incident management solutions (including custom) WIP
- Cheap to make, free to host
- Easily customize and push updates over the air with balena
- Expose your device to webhooks without the need for any backend infrastructure
- Make it excruciatingly apparent when there is an incident while on call with flashing lights and a buzzer
- Alarm automatically disables once the incident is acknowledged, or a minute has passed since activation


## Required Parts

- Raspberry Pi (Zero W in my case)
- Some sort of AC power relay or MOSFET for power switching the light and buzzer (I used a 120v AC Relay with a 120v light and buzzer)
- A balena account
