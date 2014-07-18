require 'spec_helper'

describe 'Soothe' do
  describe '/soothe' do
    it 'soothes' do
      uri = URI.parse("http://localhost:9999/soothe")
      response = Net::HTTP.get_response(uri)

      expect(response.code).to eq('200')
      json = JSON.parse(response.body)
      expect(json['soothe']).to_not be nil
      expect(json['soothe'].scan(/calmingmanatee/).count).to eq(1)
    end
  end
end