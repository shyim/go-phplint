<!DOCTYPE html>
<html lang="<?= htmlspecialchars($locale ?? 'en') ?>">
<head>
    <meta charset="utf-8">
    <title><?php echo htmlspecialchars($title ?? 'Catalog'); ?></title>
</head>
<body>
<header>
    <h1><?= htmlspecialchars($title ?? 'Catalog') ?></h1>
    <?php if (!empty($user)): ?>
        <p>Hello <?= htmlspecialchars($user['name']) ?>!</p>
    <?php else: ?>
        <p><a href="/login">Sign in</a></p>
    <?php endif; ?>
</header>

<main>
    <?php if (empty($products)): ?>
        <p class="empty">No products found.</p>
    <?php else: ?>
        <ul class="products">
            <?php foreach ($products as $index => $product): ?>
                <li class="product product--<?= $index % 2 === 0 ? 'even' : 'odd' ?>">
                    <span class="title"><?= htmlspecialchars($product['title']) ?></span>
                    <span class="price"><?php printf('%.2f', $product['price']); ?></span>
                    <?php if (!empty($product['tags'])): ?>
                        <ul class="tags">
                            <?php foreach ($product['tags'] as $tag): ?>
                                <li><?= htmlspecialchars($tag) ?></li>
                            <?php endforeach; ?>
                        </ul>
                    <?php endif; ?>
                </li>
            <?php endforeach; ?>
        </ul>
    <?php endif; ?>

    <?php
    $total = 0.0;
    foreach ($products ?? [] as $product) {
        $total += (float) $product['price'];
    }
    ?>

    <p class="total">Total: <?= number_format($total, 2) ?></p>

    <?php switch ($view ?? 'grid'):
        case 'list': ?>
            <p>List view</p>
            <?php break;
        case 'grid': ?>
            <p>Grid view</p>
            <?php break;
        default: ?>
            <p>Unknown view</p>
    <?php endswitch; ?>
</main>

<footer>
    <?php for ($page = 1; $page <= ($pages ?? 3); $page++): ?>
        <a href="?page=<?= $page ?>"><?= $page ?></a>
    <?php endfor; ?>
    <small>&copy; <?= date('Y') ?></small>
</footer>
</body>
</html>
