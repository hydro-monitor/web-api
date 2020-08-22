#!/bin/bash

 mockery --dir ./pkg/clients/db --name Client --output ./pkg/clients/db --outpkg db --filename client_mock.go --structname ClientMock
 mockery --dir ./pkg/repositories --name Repository --output ./pkg/repositories --outpkg repositories --filename repository_mock.go --structname RepositoryMock
 mockery --dir ./pkg/services --name NodeService --output ./pkg/services --inpackage --outpkg services --filename node_service_mock.go --structname NodeServiceMock
 mockery --dir ./pkg/services --name ReadingsService --output ./pkg/services --inpackage --outpkg services --filename readings_service_mock.go --structname ReadingsServiceMock
 mockery --dir ./pkg/services --name UsersService --output ./pkg/services --inpackage --outpkg services --filename users_service_mock.go --structname UsersServiceMock
