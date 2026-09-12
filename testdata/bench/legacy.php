<?php

namespace Legacy\Shop;

use InvalidArgumentException;

abstract class Model
{
    /** @var array */
    protected $attributes = array();

    public function __construct(array $attributes = array())
    {
        foreach ($attributes as $key => $value) {
            $this->set($key, $value);
        }
    }

    public function set($key, $value)
    {
        if (!is_string($key) || $key === '') {
            throw new InvalidArgumentException('invalid attribute name');
        }
        $this->attributes[$key] = $value;

        return $this;
    }

    public function get($key, $default = null)
    {
        return isset($this->attributes[$key]) ? $this->attributes[$key] : $default;
    }

    abstract public function table();
}

class Order extends Model
{
    const STATUS_OPEN = 'open';
    const STATUS_PAID = 'paid';

    public function table()
    {
        return 'orders';
    }

    public function total()
    {
        $total = 0.0;
        foreach ((array) $this->get('lines', array()) as $line) {
            $total += $line['quantity'] * $line['price'];
        }

        return round($total, 2);
    }

    public function isPaid()
    {
        return $this->get('status', self::STATUS_OPEN) === self::STATUS_PAID;
    }
}

interface OrderRepository
{
    public function findById($id);

    public function persist(Order $order);
}

class InMemoryOrderRepository implements OrderRepository
{
    private $orders = array();

    public function findById($id)
    {
        return isset($this->orders[$id]) ? $this->orders[$id] : null;
    }

    public function persist(Order $order)
    {
        $this->orders[$order->get('id')] = $order;
    }

    public function paidOrders()
    {
        return array_filter($this->orders, function (Order $order) {
            return $order->isPaid();
        });
    }
}

function format_money($amount, $currency = 'EUR')
{
    switch ($currency) {
        case 'EUR':
            $suffix = ' €';
            break;
        case 'USD':
            $suffix = ' $';
            break;
        default:
            $suffix = ' ' . $currency;
    }

    return number_format((float) $amount, 2, ',', '.') . $suffix;
}

$repository = new InMemoryOrderRepository();

for ($i = 1; $i <= 25; $i++) {
    $order = new Order(array(
        'id' => $i,
        'status' => $i % 2 === 0 ? Order::STATUS_PAID : Order::STATUS_OPEN,
        'lines' => array(
            array('quantity' => $i, 'price' => 9.95),
            array('quantity' => 2, 'price' => 19.5),
        ),
    ));
    $repository->persist($order);
}

$sum = 0.0;
foreach ($repository->paidOrders() as $order) {
    $sum += $order->total();
}

try {
    echo format_money($sum), PHP_EOL;
} catch (\Exception $exception) {
    echo $exception->getMessage(), PHP_EOL;
} finally {
    unset($repository);
}
