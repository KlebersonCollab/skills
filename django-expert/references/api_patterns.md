# Django REST Framework (DRF) Patterns

## Serializers
Sempre valide dados de entrada e transforme saídas de forma explícita.

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
Utilize ViewSets para agrupar lógica de CRUD e ações customizadas.

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
        # Lógica de compra...
        return Response({'status': 'purchased'})
```

## Auth & Permissions
- Use `IsAuthenticated` por padrão.
- Crie permissões customizadas para regras de negócio (ex: `IsOwnerOrReadOnly`).
