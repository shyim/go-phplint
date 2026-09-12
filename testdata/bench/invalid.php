<?php

declare(strict_types=1);

namespace App\Broken;

final class Basket
{
    public const LIMIT = 10;

    private array $items = [];

    public function add(string $sku, int $quantity = 1): void
    {
        if ($quantity < 1 {
            throw new \InvalidArgumentException('quantity must be positive');
        }

        $this->items[$sku] = ($this->items[$sku] ?? 0) + $quantity;
    }

    public function remove(string $sku): void
    {
        unset($this->items[$sku])
    }

    public function count(): int
    {
        return array_sum($this->items);
    }
}

function checkout(Basket $basket): void
{
    if ($basket->count() > Basket::LIMIT) {
        break;
    }

    echo 'ok', PHP_EOL;
}
