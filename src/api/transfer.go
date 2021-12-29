package api

import (
	"bytes"
	"fmt"
	"highload-wallet-api/src/config"
	"highload-wallet-api/src/jrpc"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type TransferData struct {
	TransferTasks []struct {
		DestAddr  string `json:"dest_address"`
		AmountTon string `json:"amount_ton"`
		Msg       string `json:"msg"`
	} `json:"transfer_tasks"`
}

var okres = struct {
	Status bool `json:"status"`
}{
	Status: true,
}

func Transfer(c *fiber.Ctx) error {
	var data TransferData

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(apierrs.ErrorJsonData)
	}

	var ordertxt string
	for i, tr := range data.TransferTasks {
		ordertxt += fmt.Sprintf("SEND %s %s %s", tr.DestAddr, tr.AmountTon, tr.Msg)
		if i != len(data.TransferTasks)-1 {
			ordertxt += "\n"
		}
	}

	ordfileName := path.Join(config.Cfg.TempPath.Orders, "order_"+utils.UUID()+".txt")
	if err := os.WriteFile(ordfileName, []byte(ordertxt), 0644); err != nil {
		// TODO: log error
		return c.JSON(httperrs.InternalServerError500)
	}

	bocfileName := path.Join(config.Cfg.TempPath.Bocs, "q_"+utils.UUID())

	defer func() {
		os.Remove(ordfileName)
		os.Remove(bocfileName)
	}()

	cmd := exec.Command(
		config.Cfg.Fift.Binary, // fift path
		"-s",
		config.Cfg.Contract.NewOrderFif, // path to new-order.fif
		path.Join(config.Cfg.Wallet.Path, config.Cfg.Wallet.Name), // wallet <filename-base>
		config.Cfg.Wallet.SubwalletID,                             // subwallet-id
		ordfileName,                                               // path to order file
		"--no-bounce",                                             // --no-bounce flag
		bocfileName,                                               // path to boc file
	)

	bocfileName += ".boc"

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "FIFTPATH="+config.Cfg.Fift.Path)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		fmt.Println("err")
		fmt.Println(err.Error())
		return c.JSON(httperrs.InternalServerError500)
	}
	cmd.Wait()

	// fmt.Println(stdout.String())

	s, err := jrpc.SendBocFromFile(config.Cfg.TonNet.JsonRpcURL, bocfileName)
	if err != nil {
		log.Println("ERROR SendBocFromFile: " + err.Error())
	}

	if !strings.Contains(string(s), `"ok": true`) {
		// fmt
		return c.JSON(apierrs.ErrorJsonRpc)
	}

	return c.JSON(okres)
}
