package order_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	order_model "github.com/thiago-dsd/fastfood-core-api/src/order/model"
	"github.com/thiago-dsd/fastfood-core-api/src/repository"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

// @Summary		Get all orders for the current user (paginated)
// @Description	Retrieves a paginated list of orders for the user making the request
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			user	query		order_model.QueryPaginated	true	"Query parameters for pagination and filtering orders"
// @Success			200		{array}		order_entity.Order	"List of orders"
// @Router			/api/orders [get]
// @Security		ApiKeyAuth
func GetAllOrders(c *fiber.Ctx) error {
	fmt.Println("📌 GetAllOrders chamado!")

	// Parse the query parameters from the URL
	query := new(order_model.QueryPaginated)
	if err := c.QueryParser(query); err != nil {
		fmt.Println("❌ Erro ao fazer parsing da query:", err)
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Obtém o usuário autenticado do contexto
	user, ok := c.Locals("user").(*user_entity.User)
	if !ok {
		fmt.Println("❌ Erro: Usuário não encontrado no contexto")
		return c.Status(fiber.StatusUnauthorized).JSON(common_model.NewApiError("Usuário não autenticado", nil, "handler").Send())
	}

	fmt.Println("✅ Usuário autenticado:", user.Id)

	// Busca as ordens paginadas do usuário
	orders, err := repository.GetPaginated(
		order_entity.Order{
			UserId: user.Id, // Filtra pelo ID do usuário autenticado
		},
		&query.Paginate,  // Parâmetros de paginação
		&query.DateOrder, // Ordenação por data
		&query.DateWhere, // Filtro por data
		"",
		nil,
	)
	if err != nil {
		fmt.Println("❌ Erro ao buscar ordens paginadas:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(common_model.NewApiError("Erro ao buscar pedidos", err, "repository").Send())
	}

	fmt.Printf("✅ %d pedidos encontrados (paginados)\n", len(orders))

	// Retorna os pedidos encontrados
	return c.Status(fiber.StatusOK).JSON(orders)
}
