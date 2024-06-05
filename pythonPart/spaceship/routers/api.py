import numpy as np
from fastapi import APIRouter

router = APIRouter()


@router.get('/matrix_multiply')
def matrix_multiply() -> dict:
    # Генеруємо випадкові матриці розміром 10x10
    matrix_a = np.random.rand(10, 10)
    matrix_b = np.random.rand(10, 10)
    
    # Перемножуємо матриці
    product = np.dot(matrix_a, matrix_b)
    
    # Повертаємо результат у вигляді словника
    return {
        "matrix_a": matrix_a.tolist(),
        "matrix_b": matrix_b.tolist(),
        "product": product.tolist()
    }

