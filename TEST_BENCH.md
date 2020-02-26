ntl  6:31 PM
@francesco you could certainly build a runner that executes test files in parallel. This can be done in ruby, with TestBench.





6:32
It wasn't done that way, though, because the output would be gibberish. And the environment would be more chaotic.
6:33
So, you'd need to separate the test session from the output with a queue, and then the runner would have to start multiple sessions (one for each parallel lane). Then, you'd need the output to isolate the telemetry from each session.
6:33
And that would be the beginning of the mess :wink: (edited)
6:34
Ultimately, TestBench's core ethos would never survive this process: the goal is to have minimal nonsense.
6:35
Indeed, I'd start with a bootstrap implementation.
6:35
I strongly suspect that Scott and I could have lived just fine with only TestBench's bootstrap implementation.
6:35
However, it didn't have the affordances for test abstractions that we're just now looking to start building.
6:37
I'd feel more comfortable with TestBench's bootstrap implementation than RSpec or Minitest. And it's under 300 LoC, iirc.
6:37
So, all that to say, I recommend you start with a bootstrap implementation and see how far it gets you.
