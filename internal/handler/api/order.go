package api

import (
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/pkg/auth"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderHandleApi struct {
	client       pb.OrderServiceClient
	tokenManager auth.TokenManager
}

func NewOrderHandleApi(client pb.OrderServiceClient, tokenManager auth.TokenManager, router *fiber.App) {
	orderApi := &orderHandleApi{
		client:       client,
		tokenManager: tokenManager,
	}

	routerOrder := router.Group("/api/order")

	routerOrder.Get("/hello", orderApi.handlerHello)

	routerOrder.Use(middleware.Protector())

	routerOrder.Get("/", orderApi.handleOrders)
	routerOrder.Post("/create", orderApi.handlerOrderCreate)
	routerOrder.Get("/orderbyuser", orderApi.handleOrdersByUserId)
	routerOrder.Get("/:id", orderApi.handleOrderById)

}

// @Summary Greet the user for orders
// @Description Return a greeting message for orders
// @Tags Order
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /order/hello [get]
func (h *orderHandleApi) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Order")
}

// @Summary Get all orders
// @Description Retrieve all orders
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/ [get]
func (h *orderHandleApi) handleOrders(c *fiber.Ctx) error {
	res, err := h.client.GetOrders(c.Context(), &emptypb.Empty{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get order by ID
// @Description Retrieve an order by its ID
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/{id} [get]
func (h *orderHandleApi) handleOrderById(c *fiber.Ctx) error {
	orderIdStr := c.Params("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetOrder(c.Context(), &pb.OrderRequest{
		Id: int64(orderId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get orders by user ID
// @Description Retrieve orders associated with a specific user
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/by-user [get]
func (h *orderHandleApi) handleOrdersByUserId(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetOrderUsers(c.Context(), &pb.OrderRequest{
		Id: int64(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createOrderRequest body order.CreateOrderRequest true "Request body to create a new order"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/create [post]
func (h *orderHandleApi) handlerOrderCreate(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var req order.CreateOrderRequest

	req.UserID = uint(userId)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var cartItems []*pb.CartItemRequest

	for _, item := range req.CartItems {

		cartItems = append(cartItems, &pb.CartItemRequest{
			Name:     item.Name,
			Quantity: int32(item.Quantity),
			Price:    int32(item.Price),
		})
	}

	shippingAddresReq := &pb.ShippingAddressRequest{
		Alamat:   req.ShippingAddress.Alamat,
		Kota:     req.ShippingAddress.Kota,
		Negara:   req.ShippingAddress.Negara,
		Provinsi: req.ShippingAddress.Provinsi,
	}

	res, err := h.client.CreateOrder(c.Context(), &pb.CreateOrderRequest{
		UserId:          int64(req.UserID),
		Name:            req.Name,
		Phone:           req.Phone,
		Courier:         req.Courier,
		ShippingMethod:  req.ShippingMethod,
		ShippingCost:    int32(req.ShippingCost),
		TotalProduct:    req.TotalProduct,
		TotalPrice:      req.TotalPrice,
		CartItems:       cartItems,
		ShippingAddress: shippingAddresReq,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
