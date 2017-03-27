<?php

function helloWorld($name = null)
{
  return $name ? "Hello, $name!" : 'Hello, World!';
}
