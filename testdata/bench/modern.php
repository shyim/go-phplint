<?php

declare(strict_types=1);

namespace App\Catalog;

use App\Contracts\Repository;
use App\Support\Clock;

enum Status: string
{
    case Draft = 'draft';
    case Published = 'published';
    case Archived = 'archived';

    public function label(): string
    {
        return match ($this) {
            Status::Draft => 'Draft',
            Status::Published => 'Published',
            Status::Archived => 'Archived',
        };
    }

    public function isVisible(): bool
    {
        return $this === self::Published;
    }
}

interface Identifiable
{
    public function id(): string;
}

trait Timestamps
{
    private ?\DateTimeImmutable $createdAt = null;
    private ?\DateTimeImmutable $updatedAt = null;

    public function touch(Clock $clock): static
    {
        $this->updatedAt = $clock->now();
        $this->createdAt ??= $this->updatedAt;

        return $this;
    }
}

#[\Attribute(\Attribute::TARGET_CLASS)]
final class Table
{
    public function __construct(
        public readonly string $name,
        public readonly ?string $schema = null,
    ) {
    }
}

#[Table(name: 'products', schema: 'catalog')]
final class Product implements Identifiable, \JsonSerializable
{
    use Timestamps;

    public const int MAX_TAGS = 32;

    /** @var list<string> */
    private array $tags = [];

    public function __construct(
        private readonly string $identifier,
        private string $title,
        private int|float $price = 0,
        private Status $status = Status::Draft,
    ) {
    }

    public function id(): string
    {
        return $this->identifier;
    }

    public function withTitle(string $title): self
    {
        $clone = clone $this;
        $clone->title = trim($title);

        return $clone;
    }

    public function addTag(string ...$tags): void
    {
        foreach ($tags as $tag) {
            $tag = strtolower(trim($tag));
            if ($tag === '' || in_array($tag, $this->tags, true)) {
                continue;
            }
            if (count($this->tags) >= self::MAX_TAGS) {
                throw new \OverflowException('too many tags');
            }
            $this->tags[] = $tag;
        }
    }

    public function priceWithVat(float $rate = 0.19): float
    {
        return round(((float) $this->price) * (1 + $rate), 2);
    }

    public function jsonSerialize(): array
    {
        return [
            'id' => $this->identifier,
            'title' => $this->title,
            'price' => $this->price,
            'status' => $this->status->value,
            'tags' => $this->tags,
        ];
    }
}

final class ProductRepository implements Repository
{
    /** @var array<string, Product> */
    private array $items = [];

    public function __construct(private readonly Clock $clock)
    {
    }

    public function save(Product $product): void
    {
        $this->items[$product->id()] = $product->touch($this->clock);
    }

    public function find(string $id): ?Product
    {
        return $this->items[$id] ?? null;
    }

    public function filter(callable $predicate): array
    {
        return array_values(array_filter($this->items, $predicate));
    }

    public function titles(): \Generator
    {
        foreach ($this->items as $id => $item) {
            yield $id => $item->jsonSerialize()['title'];
        }
    }
}

function summarize(iterable $products, ?callable $formatter = null): string
{
    $formatter ??= static fn (array $row): string => sprintf('%s (%s)', $row['title'], $row['status']);

    $lines = [];
    foreach ($products as $product) {
        $lines[] = $formatter($product->jsonSerialize());
    }

    return implode(PHP_EOL, $lines);
}

$repository = new ProductRepository(new Clock());
$product = (new Product('sku-1', 'Keyboard', 49.99, Status::Published))
    ->withTitle('Mechanical Keyboard');
$product->addTag('input', 'hardware', 'usb');
$repository->save($product);

$visible = $repository->filter(static fn (Product $item): bool => $item->jsonSerialize()['status'] === Status::Published->value);

echo summarize($visible), PHP_EOL;
echo $product?->priceWithVat(0.2) ?? 0.0, PHP_EOL;
