---
title: "Javascript Context"
date: 2013-08-12
draft: false
slug: js-context
city: Joinville
toc: true
tags: [javascript]
---

People make a lot of confusion about Javascript context mechanism. I don't think it's confusing at all, is just that it's different when compared with other languages we generally use.

The key to learn javascript is understanging the following:

## **The context only changes inside functions**

You might be asking something like *"but doesn't it changes in ..."* **NO**, it doesn't.

I could literally just stop writing here, but I will be kind and give you some examples.

```javascript
var global, local;

// the context will change inside the test() function declaration
function test() {
  // 'var' keyword defines a new local variable.
  var local = 1000;
  // when we don't use the 'var' keyword, it will be automatically binded
  // to the global context ('window')
  global = 100;
  return local * 2;
}

console.log(global); // undefined
console.log(local); // undefined
console.log(test()); // 2000
console.log(local); // undefined
console.log(global); // 100
```

We can also encapsulate variables with something like this:

```javascript
var global = 100;

var testFn = function() {
  // new context here
  var local = 2;

  var getLocal = function() {
    // new context here
    return local;
  };

  var setLocal = function(l) {
    // new context here
    local = l;
    return "called setLocal";
  };

  var inc = function() {
    // new context here
    local += global;
    return "called inc";
  };

  // we return an object with the functions/objects we want to expose.
  return {
    getLocal: getLocal,
    setLocal: setLocal,
    inc: inc
  };
};

var test = testFn();
console.log(test.local); // undefined
console.log(test.getLocal()); // 2
console.log(test.setLocal(10)); // called setLocal
console.log(test.getLocal()); // 10
console.log(test.inc()); // called inc
console.log(test.getLocal()); // 110
```
## Just like I said before...

The context **only** changes inside functions.

```javascript
var obj1 = {
  a: {
    b: {
      c: {
        d: this
      }
    }
  }
};

var obj2 = {
  a: {
    b: {
      c: {
        d: function() {
          // context changes here!
          return this;
        }
      }
    }
  }
};

console.log(obj1.a.b.c.d); // Window
console.log(obj2.a.b.c.d()); // Object {d: function}
```

Yep, that's it. If you want to learn more about Javascript, I strongly recommend you to read the "[JavaScript: The Good Parts](http://amzn.to/14ZmSmZ)" book, by Douglas Crockford. It will guide you through some misunderstood Javascript features in a pretty simple way.

If you have any questions, please, use the comment box below.
