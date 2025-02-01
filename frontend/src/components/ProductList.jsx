import React, { useEffect, useState } from 'react';

function ProductList() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/products');
        if (!response.ok) throw new Error('Error al cargar los productos');
        const data = await response.json();
        setProducts(data);
      } catch (error) {
        alert(`Error: ${error.message}`);
      }
    };
    fetchProducts();
  }, []);

  return (
    <div>
      <h2 className="text-xl font-semibold text-gray-700 mb-4">Lista de Productos</h2>
      {products.length === 0 ? (
        <p className="text-gray-500">No hay productos disponibles.</p>
      ) : (
        <ul className="space-y-4">
          {products.map((product) => (
            <li
              key={product.id}
              className="p-4 bg-gray-50 rounded-md shadow-sm flex justify-between items-center"
            >
              <span className="text-gray-800 font-medium">{product.name}</span>
              <span className="text-indigo-600 font-semibold">${product.price}</span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default ProductList;