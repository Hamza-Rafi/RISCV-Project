# prerequisite for instruction
li {r1}, {v1}
li {r2}, {v2}

# set pin high
lw {r3}, GPIO_OUTPUT_VAL(t0)
or {r4}, t1, {r3}
sw {r4}, GPIO_OUTPUT_VAL(t0)

# run instruction
{instruction} {r3}, {r1}, {r2}

# set pin low
not t1, t1
lw {r3}, GPIO_OUTPUT_VAL(t0)
and {r4}, t1, {r3}
sw {r4}, GPIO_OUTPUT_VAL(t0)
not t1, t1

