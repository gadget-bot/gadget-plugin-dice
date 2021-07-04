# Gadget Plugin for Dice Rolling

This is a [Gadget](https://github.com/gadget-bot/gadget) plugin to allow rolling dice.

Note that Gadget -- and therefore any plugin for Gadget -- is still very much a **work in progress**, so please don't use it in production yet (or if you do, don't complain).

## How do I use this plugin?

In your `main.go` for using Gadget, your `main()` function should instruct you bot, which is created using `gadget.Setup()` needs to be instructed to `Router.AddMentionRoutes(dice.GetMentionRoutes())`. This _usually_ looks like this:

```golang
package main

import (
  gadget "github.com/gadget-bot/gadget/core"
	dice "github.com/gadget-bot/gadget-plugin-dice"
)

func main() {
  // This is the Gadget bot
	myBot := gadget.Setup()

  // Add your custom plugins here
	myBot.Router.AddMentionRoutes(dice.GetMentionRoutes())

  // This launches your bot
	myBot.Run()
}
```

## What can this plugin do?

In your chat, you can perform the following actions:

* `dice me` - Rolls [2d6](https://en.wikipedia.org/wiki/Dice_notation) and reports the results.
  * Aliases: `roll some dice`
