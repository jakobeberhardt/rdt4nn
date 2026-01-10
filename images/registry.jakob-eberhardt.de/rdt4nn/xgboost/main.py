"""Simple XGBoost training workload.

Configured via environment variables (example for docker-compose):

  environment:
    XGB_WORKERS: "2"          # number of Python processes
    XGB_ROUNDS: "2000"        # boosting rounds per training loop
    XGB_ROWS: "100000"        # synthetic rows
    XGB_COLS: "50"            # synthetic features
    XGB_NTHREAD: "1"          # threads used by XGBoost per process
    XGB_MAX_DEPTH: "8"
    XGB_ETA: "0.1"
    XGB_SUBSAMPLE: "0.8"
    XGB_COLSAMPLE_BYTREE: "0.8"
    XGB_TREE_METHOD: "hist"

Threading controls can also be set via standard env vars:
    OMP_NUM_THREADS, OPENBLAS_NUM_THREADS, MKL_NUM_THREADS, NUMEXPR_NUM_THREADS
"""

from __future__ import annotations

import multiprocessing as mp
import os
import time

import numpy as np
import xgboost as xgb


def _env_int(name: str, default: int, *, minimum: int | None = None) -> int:
    raw = os.getenv(name)
    if raw is None or raw == "":
        value = default
    else:
        try:
            value = int(raw)
        except ValueError as exc:
            raise ValueError(f"{name} must be an int, got {raw!r}") from exc
    if minimum is not None and value < minimum:
        raise ValueError(f"{name} must be >= {minimum}, got {value}")
    return value


def _env_float(name: str, default: float, *, minimum: float | None = None) -> float:
    raw = os.getenv(name)
    if raw is None or raw == "":
        value = default
    else:
        try:
            value = float(raw)
        except ValueError as exc:
            raise ValueError(f"{name} must be a float, got {raw!r}") from exc
    if minimum is not None and value < minimum:
        raise ValueError(f"{name} must be >= {minimum}, got {value}")
    return value


def _env_str(name: str, default: str) -> str:
    value = os.getenv(name)
    return default if value is None or value == "" else value


WORKERS = _env_int("XGB_WORKERS", 1, minimum=1)
ROUNDS = _env_int("XGB_ROUNDS", 2000, minimum=1)
ROWS = _env_int("XGB_ROWS", 100_000, minimum=1)
COLS = _env_int("XGB_COLS", 50, minimum=1)

NTHREAD = _env_int("XGB_NTHREAD", 1, minimum=1)
MAX_DEPTH = _env_int("XGB_MAX_DEPTH", 8, minimum=1)
ETA = _env_float("XGB_ETA", 0.1, minimum=0.0)
SUBSAMPLE = _env_float("XGB_SUBSAMPLE", 0.8, minimum=0.0)
COLSAMPLE_BYTREE = _env_float("XGB_COLSAMPLE_BYTREE", 0.8, minimum=0.0)

TREE_METHOD = _env_str("XGB_TREE_METHOD", "hist")
OBJECTIVE = _env_str("XGB_OBJECTIVE", "binary:logistic")
EVAL_METRIC = _env_str("XGB_EVAL_METRIC", "logloss")
SEED = _env_int("XGB_SEED", 435435, minimum=0)


def worker(rank: int) -> None:
    rng = np.random.default_rng(SEED + rank)
    X = rng.normal(size=(ROWS, COLS)).astype(np.float32)
    y = (X[:, 0] + 0.1 * X[:, 1] > 0).astype(np.float32)
    dtrain = xgb.DMatrix(X, label=y)

    params: dict[str, object] = {
        "objective": OBJECTIVE,
        "eval_metric": EVAL_METRIC,
        "tree_method": TREE_METHOD,
        "max_depth": MAX_DEPTH,
        "eta": ETA,
        "subsample": SUBSAMPLE,
        "colsample_bytree": COLSAMPLE_BYTREE,
        "nthread": NTHREAD,
        "seed": SEED + rank,
    }

    while True:
        _ = xgb.train(params, dtrain, num_boost_round=ROUNDS)


def main() -> None:
    print(
        "xgboost config: "
        f"workers={WORKERS} rounds={ROUNDS} rows={ROWS} cols={COLS} nthread={NTHREAD} "
        f"tree_method={TREE_METHOD} max_depth={MAX_DEPTH} eta={ETA} subsample={SUBSAMPLE} "
        f"colsample_bytree={COLSAMPLE_BYTREE} objective={OBJECTIVE} eval_metric={EVAL_METRIC} seed={SEED}",
        flush=True,
    )

    try:
        mp.set_start_method("fork")
    except RuntimeError:
        pass

    procs: list[mp.Process] = []
    for rank in range(WORKERS):
        proc = mp.Process(target=worker, args=(rank,), daemon=False)
        proc.start()
        procs.append(proc)

    while True:
        time.sleep(3600)


if __name__ == "__main__":
    main()