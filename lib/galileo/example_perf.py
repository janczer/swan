import ga
from perf_counters import Perf
from shell import Shell
from taskset import Taskset


class PerfExperiment(ga.Experiment):
    def __init__(self):
        ga.Experiment.__init__(self)

        def baseline(configuration):
            Shell([Perf(Taskset(["0"], "echo foobar"))])
            Shell([Perf(Taskset(["0"], "sleep 1"))])
            Shell([Perf(Taskset(["0"], "echo barbaz"))])

        def experiment(configuration):
            Shell([Perf(Taskset(["0"], "sleep 1"))])

        self.add_phase("baseline", baseline)
        self.add_phase("experiment", experiment)


def main():
    s = PerfExperiment()
    s.run()


if __name__ == "__main__":
    main()