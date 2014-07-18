require 'spec_helper'

describe 'Kittens' do
  describe '/kittens/random' do
    it 'returns a random kitten' do
      uri = URI.parse("http://localhost:9999/kittens/random")
      response = Net::HTTP.get_response(uri)

      expect(response.code).to eq('200')
      json = JSON.parse(response.body)
      expect(json['kitten']).to_not be nil
      expect(json['kitten'].scan(/placekitten/).count).to eq(1)
    end
  end

  describe '/kittens/bomb/:count' do
    it 'returns :count kittens' do
      uri = URI.parse("http://localhost:9999/kittens/bomb/6")
      response = Net::HTTP.get_response(uri)

      expect(response.code).to eq('200')
      json = JSON.parse(response.body)
      expect(json['kittens'].length).to eq(6)
      expect(json['kittens'][0].scan(/placekitten/).count).to eq(1)
      expect(json['kittens'][5].scan(/placekitten/).count).to eq(1)
    end
  end
end