<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class MerchantController extends Controller
{
    public function getMerchant()
    {
        // Hit Service toko
        $response = Http::get("http://localhost:8081/get-detail-toko");
        $data = $response->json();
        $merchantDetail = [
          "nama" => "Jabutech",
          "nama_toko" => $data['nama'],
            "jumlah_product" => $data['jumlah_product']
        ];

        return response()->json($merchantDetail, 200);
    }
}
