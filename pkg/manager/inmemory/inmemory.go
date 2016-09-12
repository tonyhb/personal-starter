package inmemory

type MemMgr struct {
	*userManager
	*accountManager
}

func New() *MemMgr {
	return &MemMgr{
		NewUserManager(),
		NewAccountManager(),
	}
}
