package order

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
	"shop/app/shared"
)

var service = new(orderService)

type orderService struct {
}

func (s *orderService) Add(r *ghttp.Request, req *AddOrderReq) (res sql.Result, err error) {
	req.OrderInfo.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	req.OrderInfo.Number = shared.GetOrderNum()
	//生成主订单
	lastInsertId, err := dao.OrderInfo.Ctx(r.GetCtx()).InsertAndGetId(req.OrderInfo)
	if err != nil {
		return nil, err
	}
	//生成商品订单
	for _, info := range req.OrderGoodsInfos {
		info.OrderId = gconv.Int(lastInsertId)
		_, err := dao.OrderGoodsInfo.Ctx(r.GetCtx()).Insert(info)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (s *orderService) Update(r *ghttp.Request, req *UpdateOrderReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.OrderInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *orderService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *orderService) List(r *ghttp.Request, req *PageListReq) (res ListOrderRes, err error) {
	whereCondition := g.Map{}
	if req.Status == 0 {
		whereCondition = g.Map{
			dao.OrderInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
		}
	} else {
		whereCondition = g.Map{
			dao.OrderInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
			dao.OrderInfo.Columns.Status: req.Status,
		}
	}
	count, err := dao.OrderInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.OrderInfo.Ctx(r.GetCtx()).With(OrderGoodsInfo{}).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

//同类商品推荐
//func (s *orderService) Category(ctx context.Context, req *CategoryPageListReq) (res ListOrderRes, err error) {
//	//获取商品的分类
//	currentOrder := model.OrderInfo{}
//	err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Scan(&currentOrder)
//	if err != nil {
//		return ListOrderRes{}, err
//	}
//
//	whereLevelCondition := g.Map{
//		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{currentOrder.Level1CategoryId, currentOrder.Level2CategoryId, currentOrder.Level3CategoryId},
//	}
//	whereIdCondition := g.Map{
//		"id!=": req.Id,
//	}
//	count, err := dao.OrderInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//	err = dao.OrderInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Page(req.Page, req.Limit).Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//func (s *orderService) List(ctx context.Context, req *PageListReq) (res ListOrderRes, err error) {
//	whereCondition := g.Map{}
//	if req.Keyword != "" && req.CategoryId != 0 {
//		whereCondition = g.Map{
//			"name like": "%" + req.Keyword + "%",
//			"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{req.CategoryId, req.CategoryId, req.CategoryId},
//		}
//	} else if req.Keyword != "" {
//		whereCondition = g.Map{
//			"name like": "%" + req.Keyword + "%",
//		}
//	} else if req.CategoryId != 0 {
//		whereCondition = g.Map{
//			"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{req.CategoryId, req.CategoryId, req.CategoryId},
//		}
//	} else {
//		whereCondition = g.Map{}
//	}
//
//	//获取数量
//	count, err := dao.OrderInfo.Ctx(ctx).
//		Where(whereCondition).
//		Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//
//	//获取值
//	//排序规则
//	sortCondition := packSort(req)
//	err = dao.OrderInfo.Ctx(ctx).
//		Where(whereCondition).
//		Page(req.Page, req.Limit).
//		Order(sortCondition).
//		Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//封装排序方法
//func packSort(req *SearchPageListReq) (sortCondition string) {
//	//排序规则
//	sortCondition = dao.OrderInfo.Columns.CreatedAt + " ASC" //id升序
//	if req.Sort == "recent" {                               //最近上架
//		sortCondition = dao.OrderInfo.Columns.CreatedAt + " DESC" //创建时间倒序
//	} else if req.Sort == "sale" {
//		sortCondition = dao.OrderInfo.Columns.Sale + " DESC" //销量倒序
//	} else if req.Sort == "price_up" {
//		sortCondition = dao.OrderInfo.Columns.Price + " ASC" //价格升序
//	} else if req.Sort == "price_down" {
//		sortCondition = dao.OrderInfo.Columns.Price + " DESC" //价格降序
//	}
//	return
//}

func (s *orderService) Detail(ctx context.Context, req *DetailReq) (res ListOrderSql, err error) {
	err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}