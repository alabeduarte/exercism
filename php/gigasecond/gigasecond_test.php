<?php

require "gigasecond.php";

class GigasecondTest extends PHPUnit\Framework\TestCase
{

    public function dateSetup($date)
    {
        $UTC = new DateTimeZone("UTC");
        $date = new DateTime($date, $UTC);
        return $date;
    }

    public function testDate1()
    {
        $date = GigasecondTest::dateSetup("2011-04-25");
        $gs = from($date);

        $this->assertSame("2043-01-01 01:46:40", $gs->format("Y-m-d H:i:s"));
    }

    public function testDate2()
    {
        $date = GigasecondTest::dateSetup("1977-06-13");
        $gs = from($date);

        $this->assertSame("2009-02-19 01:46:40", $gs->format("Y-m-d H:i:s"));
    }

    public function testPreUnixEpoch()
    {
        $date = GigasecondTest::dateSetup("1959-7-19");
        $gs = from($date);

        $this->assertSame("1991-03-27 01:46:40", $gs->format("Y-m-d H:i:s"));
    }

    public function testDateWithTime1()
    {
        $date = GigasecondTest::dateSetup("2015-01-24 22:00:00");
        $gs = from($date);

        $this->assertSame("2046-10-02 23:46:40", $gs->format("Y-m-d H:i:s"));
    }

    public function testDateWithTime2()
    {
        $date = GigasecondTest::dateSetup("2015-01-24 23:59:59");
        $gs = from($date);

        $this->assertSame("2046-10-03 01:46:39", $gs->format("Y-m-d H:i:s"));
    }

    public function testNoMutation()
    {
        $date = GigasecondTest::dateSetup("2015-01-24");
        $gs = from($date);

        $this->assertNotEquals($date, $gs);
    }

    public function testYourself()
    {
        $my_birthday = GigasecondTest::dateSetup("1988-01-23 07:00:00");
        $gs = from($my_birthday);

        $this->assertSame("2019-10-01 08:46:40", $gs->format("Y-m-d H:i:s"));
    }
}
