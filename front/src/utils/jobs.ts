const _lastRunTimeout = Symbol('last run timeout');
type _LastRunJob = {
    timeout: NodeJS.Timer,
    afterRun: Set<Function>,
};

type _LastRunJobs = {
    [jobName: string]: _LastRunJob,
};

/**
 * Call multiple times, only last `jobf` after timeout will be called
 * Timeout resets on every call
 * `target` may be any object with fields
 */
export function runLastTimeout(
    target: object,
    name: string,
    jobf: Function,
    timeout: number = 120,
    afterRun?: Function
) {
    // @ts-ignore
    if (!target[_lastRunTimeout]) target[_lastRunTimeout] = {};

    // @ts-ignore
    const jobs: _LastRunJobs = target[_lastRunTimeout];

    let job: _LastRunJob;
    if (name in jobs) {
        job = jobs[name];
        clearTimeout(job.timeout);
    } else {
        job = {
            timeout: undefined!,
            afterRun: new Set(),
        };
        jobs[name] = job;
    }

    job.timeout = setTimeout(() => {
        jobf();
        job.afterRun.forEach(f => f());
        job.afterRun.clear();
    }, timeout) as any;

    if (afterRun) job.afterRun.add(afterRun);
}
