## Lightweight background worker package
<img src="img/worker.jpg" width="120" height="100">

Package provides basic functionality for background concurrent worker.

### General information

Class `Worker` represents a worker that runs specified function in the background of your application.
It is designed to be flexible, customizable and respective towards higher-level `Context`.

For now this package has no dependencies, and I intend to keep it this way.

### How to use

To create new `Worker` constructor func `New` should be used. It allowes to use multiple options:
- `WithDelay` sets time delay between finishing N-th run and starting (N+1)-th run;
- `WithOnErrDelay` sets delay that is applied if worker's run finished with error;
- `WithCurrentTimeProvider` allows to specify a function which returns current time which is extremely useful in tests;
- `WithTaskRunTimeInterval` sets time interval (hours, minutes, seconds) in which worker is allowed to run his task.

