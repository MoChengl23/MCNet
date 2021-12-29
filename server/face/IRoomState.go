package face

type IRoomState interface {
	Enter()
	Exit()
	Update()
}
