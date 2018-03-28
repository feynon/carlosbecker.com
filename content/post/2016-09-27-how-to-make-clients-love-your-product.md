---
date: 2016-09-27T00:00:00Z
slug: how-to-make-clients-love-your-product
title: How to make people love your product
title:
- product
---

I have seen **a lot** of posts like this subject, but almost all of them were
about stuff you should do and almost none of them reflected how I
truly feel, how I want a product to be so I can "love" it.

If I would make a list, it would look more or less like this:

### 1. No bullshit

- don't send me a hundred emails;
- don't use [dark design patterns](http://darkpatterns.org/);
- for the christ's sake, no fucking popups;
- if I'm paying, don't show me ads;
- if I'm not paying, it's because I didn't like the product, stop sending me
emails;
- don't call me unless I ask you to do so;
- if you need to send emails to remember the user that your product exists,
your product sucks;

### 2. Security, please

- [the only accepted way of validating emails is to send a confirmation email](https://hackernoon.com/the-100-correct-way-to-validate-email-addresses-7c4818f24643).
Just get over it and get it done. If you block email tags I will say bad words
about your mother and close the browser tab shouting "fuck that";
- don't implement stupid password validations. According to
[How Secure Is My Password](https://howsecureismypassword.net/),
`aSd_123` takes 7 minutes to crack up (and would possibly pass your validations).
`asd123` is instant (may or may not pass your validations).
`dont-set-password-size-limits-dumbass` takes 373 DUODECILLION YEARS and most
sites would complain because it doesn't have any numbers or capital letters.
Just delete that code and drop the constraints already;
- MFA/2FA is required depending on the information the product holds;
- one of the first things I do when I sign up to something is recover the
password. If you send me the password in plain text you will get a not very
educated but very educational email from me - and I will never, ever,
recommend your product to anyone.

### 3. No dumbassery

- don't sell things you didnt't yet have and/or don't yet work as expected;
- don't break the fucking browser features. I want to be able to press the
back button. And it should work. Added some JS magic? Add it the right way;
- don't re-do browser components making them less user friendly than the
original ones (like scroll bars for numbers in which I can't type in).

### 4. The less interaction the better

- if I would like to talk to people, I'll go to a bar. Don't fucking call me
unless I ask you to;
- I know the email is automated, the sender is `no-reply@product.com`,
don't ask me rhetorical question;
- have a clear and updated documentation. The last thing I want to
do is to contact a human while solving a software-related issue;
- have an API, so I can automate stuff;
- don't send me useless emails.

-----

These are my 4 pillars for a product I might love.

If you take a closer look at this list, most of these itens are things you
shouldn't do. People you dont't need to hire. Code that don't exist
(and therefore don't break nor get slow).

But, for some reason, some PO's seem to love some of those things,
and I am just a unknown grumpy developer. What the fuck do I know, right?

