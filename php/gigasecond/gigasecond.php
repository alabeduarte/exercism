<?php

function from($date) {
  $gigasecond = new DateInterval('PT'.pow(10, 9).'S');

  return date_add(clone $date, $gigasecond);
}
