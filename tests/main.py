import unittest

import grpc

from user_pb2 import CreateOrReturnUserRequest, ChangeLangRequest
from user_pb2_grpc import UsersStub
from totp_pb2 import AddToTPRequest, FindAllToTPRequest, RemoveToTPRequest
from totp_pb2_grpc import ToTPsStub


class TestGrpc(unittest.TestCase):
    def setUp(self):
        channel = grpc.insecure_channel('localhost:50051')
        self.users_stub = UsersStub(channel)
        self.totps_stub = ToTPsStub(channel)

    def test_create_or_return_user(self):
        request = CreateOrReturnUserRequest(userId=123)
        response = self.users_stub.CreateOrReturnUser(request)
        self.assertIsNotNone(response.response)

    def test_change_lang(self):
        request = ChangeLangRequest(userId=123, lang="en")
        response = self.users_stub.ChangeLang(request)
        self.assertTrue(response.status)

    def test_add_totp(self):
        request = AddToTPRequest(userId=123, totp="123456", name="MyTOTP")
        response = self.totps_stub.AddToTP(request)
        self.assertTrue(response.status)

    def test_find_all_totp(self):
        request = FindAllToTPRequest(userId=123)
        response = self.totps_stub.FindAllToTP(request)
        self.assertTrue(len(response.response) > 0)

    def test_remove_totp(self):
        request = RemoveToTPRequest(userId=123, totp="123456")
        response = self.totps_stub.RemoveToTP(request)
        self.assertTrue(response.status)

if __name__ == '__main__':
    unittest.main()
