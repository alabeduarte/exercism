<?php

function zip($arrayA, $arrayB) {
  return array_map(null, $arrayA, $arrayB);
};

function distance($a, $b)
{
  if (strlen($a) != strlen($b)) {
    throw new InvalidArgumentException('DNA strands must be of equal length.');
  }

  return count(
    array_filter(zip(str_split($a), str_split($b)), function($pair) {
      return $pair[0] != $pair[1];
    })
  );
};
