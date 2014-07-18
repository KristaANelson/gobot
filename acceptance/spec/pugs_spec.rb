require 'spec_helper'

describe 'Pugs' do
  describe '/pugs/random' do
    it 'returns a single pug' do
      uri = URI.parse("http://localhost:9999/pugs/random")
      response = Net::HTTP.get_response(uri)

      expect(response.code).to eq('200')
      json = JSON.parse(response.body)
      expect(json['pug']).to_not be nil
      expect(json['pug'].length).to be > 10
    end
  end

  describe '/pugs/bomb/:count' do
    it 'returns :count pugs' do
      uri = URI.parse("http://localhost:9999/pugs/bomb/4")
      response = Net::HTTP.get_response(uri)

      expect(response.code).to eq('200')
      json = JSON.parse(response.body)
      expect(json['pugs'].length).to eq(4)
      expect(json['pugs'][0].length).to be > 10
      expect(json['pugs'][3].length).to be > 10
    end
  end
end