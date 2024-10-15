---
title: "JHiccup Jitter¶"
weight: 2
---

To be able to easily detect long GC pauses or hardware overload we’ve
added a StatModule similar to
[External Link: jHiccup](https://github.com/giltene/jHiccup). The idea is to compute
the time spent to sleep 1 millisecond and allocate an object. If this
time varies it is caused by GC, thread scheduling or machine over
provisioning…

The KPI is reported to hubble as _JHiccup jitter in Nanos_ here’s an
example of observation:

![image](http://monitoring/../_images/jhiccup_hubble.png){height=”400px”}

I ran this locally on a busy machine so you can see the 90th percentile
is around 20 microseconds. I then called _docker pause_ to simulate
pauses, this created the two big spikes in maximum of about 4 and 8
seconds. [External Link: JHiccup](https://github.com/giltene/jHiccup) has the advantage
of detecting all kind of pauses even if the most common one will be GC.

You can disable reporting with
_pie.queue.metrics.jhiccup.disabled=true_. It is possible to extra the
entire [External Link: HdrHistogram](https://hdrhistogram.github.io/HdrHistogram/) to a file by setting
_pie.queue.metrics.jhiccup.logFileName_. Beware that there is currently
no rolling so you should therefore not use this in production.

If you write this file you can use the [External Link: Excel document in the jHiccup
repo](https://github.com/giltene/jHiccup/blob/master/jHiccupPlotter.xls)
to generate reports and more precise graphs. This can be useful when
running performance testing for example.
