# Django REST Framework (DRF) Patterns

## Serializers
Always validate input data and transform outputs explicitly.

```python
class ProductSerializer(serializers.ModelSerializer):
    category_name = serializers.CharField(source='category.name', read_only=True)
    
    class Meta:
        model = Product
        fields = ['id', 'name', 'price', 'category_name']
        read_only_fields = ['id']

    def validate_price(self, value):
        if value < 0:
            raise serializers.ValidationError("Price cannot be negative.")
        return value
```

## ViewSets
Use ViewSets to group CRUD logic and customized actions.

```python
class ProductViewSet(viewsets.ModelViewSet):
    queryset = Product.objects.all()
    serializer_class = ProductSerializer
    
    def get_serializer_class(self):
        if self.action == 'create':
            return ProductCreateSerializer
        return self.serializer_class

    @action(detail=True, methods=['post'])
    def purchase(self, request, pk=None):
        product = self.get_object()
        # Purchase logic...
        return Response({'status': 'purchased'})
```

## Auth & Permissions
- Use `IsAuthenticated` by default.
- Create custom permissions for business rules (e.g., `IsOwnerOrReadOnly`).
