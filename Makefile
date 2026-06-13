.PHONY: swagger

# 从接口注释生成 Swagger 文档
# 依赖: go install github.com/swaggo/swag/cmd/swag@latest
swagger:
	swag init
