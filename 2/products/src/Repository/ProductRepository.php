<?php

namespace App\Repository;

use App\Entity\Product;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends ServiceEntityRepository<Product>
 */
class ProductRepository extends ServiceEntityRepository {
	public function __construct(ManagerRegistry $registry) {
		parent::__construct($registry, Product::class);
	}

	public function getAll(): array {
		return $this->createQueryBuilder('p')
			->getQuery()
			->getResult()
		;
	}

	public function getById($id): array {
		return $this->createQueryBuilder('p')
			->andWhere('p.id = :id')
			->setParameter('id', $id)
			->getQuery()
			->getResult()
		;
	}
}
