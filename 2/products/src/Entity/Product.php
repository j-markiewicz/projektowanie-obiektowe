<?php

namespace App\Entity;

use App\Repository\ProductRepository;
use Doctrine\ORM\Mapping as ORM;
use JsonSerializable;

#[ORM\Entity(repositoryClass: ProductRepository::class)]
class Product implements JsonSerializable {
	#[ORM\Id]
	#[ORM\GeneratedValue]
	#[ORM\Column]
	private ?int $id = null;

	#[ORM\Column(length: 255)]
	private ?string $name = null;

	#[ORM\Column(length: 1023)]
	private ?string $description = null;

	#[ORM\Column]
	private ?int $price = null;

	public function getId(): ?int {
		return $this->id;
	}

	public function setId(int $id): static {
		$this->id = $id;

		return $this;
	}

	public function getName(): ?string {
		return $this->name;
	}

	public function setName(string $name): static {
		$this->name = $name;

		return $this;
	}

	public function getDescription(): ?string {
		return $this->description;
	}

	public function setDescription(string $description): static {
		$this->description = $description;

		return $this;
	}

	public function getPrice(): ?int {
		return $this->price;
	}

	public function setPrice(int $price): static {
		$this->price = $price;

		return $this;
	}

	public function jsonSerialize(): mixed {
		return [
			"id" => $this->getId(),
			"name" => $this->getName(),
			"description" => $this->getDescription(),
			"price" => $this->getPrice(),
		];
	}
}
