# go-ddd


# todo
- [ ] readme
- [ ] telemetry
- [ ] command validation
- [x] result pattern
- [ ] define events
- [ ] outbox
- [x] db migration run same migration on test db
- [ ] tests
  - [x] unit
  - [ ] service
- [ ] pipeline
  - [ ] ci
- [ ] infrastructure
  - [ ] helm


go install github.com/vektra/mockery/v2@latest <br>
mockery --name=OrderRepository --output=mocks <br>
go-ddd % go install github.com/fatih/gomodifytags