#!/bin/sh
gunicorn -b 0.0.0.0:8000 -w 1 -t 60 app:app --reload 
