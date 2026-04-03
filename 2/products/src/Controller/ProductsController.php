<?php

namespace App\Controller;

use App\Entity\Product;
use App\Repository\ProductRepository;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

final class ProductsController extends AbstractController {
	#[Route("/products", "list_products", methods: ["GET"])]
	public function listProduct(ProductRepository $productRepository): JsonResponse {
		return $this->json($productRepository->getAll());
	}

	#[Route("/products/{id}", "get_product", methods: ["GET"])]
	public function getProduct(Product $product): JsonResponse {
		return $this->json($product);
	}

	#[Route("/products", "create_product", methods: ["POST"])]
	public function createProduct(Request $request, EntityManagerInterface $entityManager): JsonResponse {
		$requestBody = json_decode($request->getContent(), true);

		$product = new Product();
		$product->setName($requestBody["name"]);
		$product->setDescription($requestBody["description"]);
		$product->setPrice(intval($requestBody["price"]));

		$entityManager->persist($product);
		$entityManager->flush();

		return $this->json($product, status: Response::HTTP_CREATED);
	}

	#[Route("/products/{id}", "delete_product", methods: ["DELETE"])]
	public function deleteProduct(Product $product, EntityManagerInterface $entityManager) {
		$entityManager->remove($product);
		$entityManager->flush();

		return $this->json(null, Response::HTTP_NO_CONTENT);
	}

	#[Route("/products/{id}", "update_product", methods: ["PATCH", "PUT"])]
	public function updateProduct(Product $product, Request $request, EntityManagerInterface $entityManager) {
		$requestBody = json_decode($request->getContent(), true);

		$product->setName($requestBody["name"]);
		$product->setDescription($requestBody["description"]);
		$product->setPrice(intval($requestBody["price"]));

		$entityManager->persist($product);
		$entityManager->flush();

		return $this->json($product);
	}
}
