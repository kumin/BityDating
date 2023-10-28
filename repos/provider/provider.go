package provider

import (
	"github.com/google/wire"
	"github.com/kumin/AndPadDating/infras"
	"github.com/kumin/AndPadDating/repos"
	"github.com/kumin/AndPadDating/repos/mysql"
)

var MysqlGraphSet = wire.NewSet(
	infras.InfaGraphSet,
	mysql.NewUserMysqlRepo,
	wire.Bind(new(repos.UserRepo), new(*mysql.UserMysqlRepo)),
	mysql.NewMatchingMysqlRepo,
	wire.Bind(new(repos.MatchingRepo), new(*mysql.MatchingMysqlRepo)),
	mysql.NewFeedMysqlRepo,
	wire.Bind(new(repos.FeedRepo), new(*mysql.FeedMysqlRepo)),
)