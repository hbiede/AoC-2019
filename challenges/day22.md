## Day 22: Slam Shuffle

 There isn't much to do while you wait for the droids to repair your ship.

At least you're drifting in the right direction.

You decide to practice a new card shuffle you've been working on.

Digging through the ship's storage, you find a deck of space cards ! Just like any deck of space cards , there are 10007 cards in the deck numbered ` 0 ` through ` 10006 `.

The deck must be new - they're still in factory order , with ` 0 ` on the top, then ` 1 ` , then ` 2 ` , and so on, all the way through to ` 10006 ` on the bottom.

You've been practicing three different techniques that you use while shuffling.

Suppose you have a deck of only 10 cards (numbered ` 0 ` through ` 9 ` ): To ` deal into new stack ` , create a new stack of cards by dealing the top card of the deck onto the top of the new stack repeatedly until you run out of cards: ` Top Bottom0 1 2 3 4 5 6 7 8 9 Your deck New stack 1 2 3 4 5 6 7 8 9 Your deck 0 New stack 2 3 4 5 6 7 8 9 Your deck 1 0 New stack 3 4 5 6 7 8 9 Your deck 2 1 0 New stackSeveral steps later... 9 Your deck 8 7 6 5 4 3 2 1 0 New stack Your deck9 8 7 6 5 4 3 2 1 0 New stack ` Finally, pick up the new stack you've just created and use it as the deck for the next technique.

To ` cut N ` cards , take the top ` N ` cards off the top of the deck and move them as a single unit to the bottom of the deck, retaining their order.

For example, to ` cut 3 ` : ` Top Bottom0 1 2 3 4 5 6 7 8 9 Your deck 3 4 5 6 7 8 9 Your deck0 1 2 Cut cards3 4 5 6 7 8 9 Your deck 0 1 2 Cut cards3 4 5 6 7 8 9 0 1 2 Your deck ` You've also been getting pretty good at a version of this technique where ` N ` is negative! In that case, cut (the absolute value of) ` N ` cards from the bottom of the deck onto the top.

For example, to ` cut -4 ` : ` Top Bottom0 1 2 3 4 5 6 7 8 9 Your deck0 1 2 3 4 5 Your deck 6 7 8 9 Cut cards 0 1 2 3 4 5 Your deck6 7 8 9 Cut cards6 7 8 9 0 1 2 3 4 5 Your deck ` To ` deal with increment N ` , start by clearing enough space on your table to lay out all of the cards individually in a long line.

Deal the top card into the leftmost position.

Then, move ` N ` positions to the right and deal the next card there.

If you would move into a position past the end of the space on your table, wrap around and keep counting from the leftmost card again.

Continue this process until you run out of cards.

For example, to ` deal with increment 3 ` : ` 0 1 2 3 4 5 6 7 8 9 Your deck..........

Space on table^ Current positionDeal the top card to the current position: 1 2 3 4 5 6 7 8 9 Your deck0.........

Space on table^ Current positionMove the current position right 3: 1 2 3 4 5 6 7 8 9 Your deck0.........

Space on table ^ Current positionDeal the top card: 2 3 4 5 6 7 8 9 Your deck0.. 1......

Space on table ^ Current positionMove right 3 and deal: 3 4 5 6 7 8 9 Your deck0.. 1.. 2...

Space on table ^ Current positionMove right 3 and deal: 4 5 6 7 8 9 Your deck0.. 1.. 2.. 3 Space on table ^ Current positionMove right 3, wrapping around, and deal: 5 6 7 8 9 Your deck0. 4 1.. 2.. 3 Space on table ^ Current positionAnd so on:0 7 4 1 8 5 2 9 6 3 Space on table ` Positions on the table which already contain cards are still counted; they're not skipped.

Of course, this technique is carefully designed so it will never put two cards in the same position or leave a position empty.

Finally, collect the cards on the table so that the leftmost card ends up at the top of your deck, the card to its right ends up just below the top card, and so on, until the rightmost card ends up at the bottom of the deck.

The complete shuffle process (your puzzle input) consists of applying many of these techniques.

Here are some examples that combine techniques; they all start with a factory order deck of 10 cards: ` deal with increment 7deal into new stackdeal into new stackResult: 0 3 6 9 2 5 8 1 4 7 ` ` cut 6deal with increment 7deal into new stackResult: 3 0 7 4 1 8 5 2 9 6 ` ` deal with increment 7deal with increment 9cut -2Result: 6 3 0 7 4 1 8 5 2 9 ` ` deal into new stackcut -2deal with increment 7cut 8cut -4deal with increment 7cut 3deal with increment 9deal with increment 3cut -1Result: 9 2 5 8 1 4 7 0 3 6 ` Positions within the deck count from ` 0 ` at the top, then ` 1 ` for the card immediately below the top card, and so on to the bottom. (That is, cards start in the position matching their number.) After shuffling your factory order deck of 10007 cards, what is the position of card ` 2019 ` ? 
