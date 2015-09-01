#!/bin/sh
gunicorn -b 0.0.0.0:8000 -k aiohttp.worker.GunicornWebWorker -w 1 -t 60 app:app --reload 
