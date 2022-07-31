# GoLangEnumFlag

    type PlayerStatus uint64

    const (
      PlayerStatus_None    PlayerStatus = 0
      PlayerStatus_NewUser PlayerStatus = (1 << 0)
      PlayerStatus_Banned  PlayerStatus = (1 << 1)

      PlayerStatus_Max PlayerStatus = (1 << 63)
    )

    func main() {
      playerStatus := EnumFlag.NewEnumFlag[PlayerStatus]()

      //playerStatus.Store(3) // PlayerStatus_NewUser(1) + PlayerStatus_Banned(2) => 3


      playerStatus.Combine(PlayerStatus_NewUser)
      //playerStatus.Combine(PlayerStatus_Banned)
      playerStatus.Combine(PlayerStatus_Max)

      log.Println("HasFlag(Banned):", playerStatus.HasFlag(PlayerStatus_Banned))
      log.Println("HasFlag(NewUser):", playerStatus.HasFlag(PlayerStatus_NewUser))
      log.Println("HasFlag(Max):", playerStatus.HasFlag(PlayerStatus_Max))
      log.Println("Value:", playerStatus.Load())
    }
![image](https://user-images.githubusercontent.com/59788044/182005194-257ecc42-b458-4d82-8423-3c0bc4675345.png)
