---
title: "Exponential Backoff with Java 8"
date: 2014-10-26
draft: false
slug: exponential-backoff-java8
city: Joinville
toc: true
tags: [java]
---

> Exponential backoff is an algorithm that uses feedback to multiplicatively decrease the rate of some process, in order to gradually find an acceptable rate. 
> —- Wikipedia

I recently used this strategy in work to deal with another service that we need to integrate. Sometimes, the service will just refuse the connection, without any reason. If I keep pushing, it will, someday, accept it.

So, I used Java 8 Functional Interfaces to implement this in a not-so-ugly way, also using a Fibonacci's Sequence to increment the wait time:

### The `ExponentialBackOffFunction` Functional Interface:

```java
import java.rmi.RemoteException;

@FunctionalInterface
public interface ExponentialBackOffFunction<T> {
	T execute();
}
```

### The `ExponentialBackOff` main class:

```java
import static java.util.Arrays.asList;

import java.net.SocketTimeoutException;
import java.util.List;

import javax.net.ssl.SSLHandshakeException;

import lombok.extern.log4j.Log4j;

@Log4j
public final class ExponentialBackOff {
	private static final int[] FIBONACCI = new int[] { 1, 1, 2, 3, 5, 8, 13 };
	private static final List<Class<? extends Exception>> EXPECTED_COMMUNICATION_ERRORS = asList(
			SSLHandshakeException.class, SocketTimeoutException.class );

	private ExponentialBackOff() {

	}

	public static <T> T execute(ExponentialBackOffFunction<T> fn) {
    for (int attempt = 0; attempt < FIBONACCI.length; attempt++) {
			try {
				return fn.execute();
			} catch (Exception e) {
				handleFailure( attempt, e );
			}
		}
		throw new RuntimeException( "Failed to communicate." );
	}

	private static void handleFailure(int attempt, RemoteException e) {
		if (e.getCause() != null && !EXPECTED_COMMUNICATION_ERRORS.contains( e.getCause().getClass() ))
			throw new RuntimeException( e );
		doWait( attempt );
	}

	private static void doWait(int attempt) {
		try {
			Thread.sleep( FIBONACCI[attempt] * 1000 );
		} catch (InterruptedException e) {
			throw new RuntimeException( e );
		}
	}
}
```

### Usage:

```java
ExponentialBackoff.execute( () -> work() );
```

This will try to execute the `work` method incrementing the time between each call that fail with an expected error.
