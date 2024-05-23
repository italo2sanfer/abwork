run:
	@uvicorn workout_api.main:app --reload

run_2:
	@fastapi run api/main.py --port 8000

create-migrations:
	@PYTHONPATH=$PYTHONPATH:$(pwd) alembic revision --autogenerate -m $(d)

run-migrations:
	@PYTHONPATH=$PYTHONPATH:$(pwd) alembic upgrade head