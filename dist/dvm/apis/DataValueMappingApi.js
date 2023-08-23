"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud DataValueMapping Service API
 * APIs which offer external systems to config and use values different than what are configured in opera<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.GetConvertedValuesConversionCodeEnum = exports.DataValueMappingApi = void 0;
const runtime = __importStar(require("../runtime"));
const models_1 = require("../models");
/**
 *
 */
class DataValueMappingApi extends runtime.BaseAPI {
    /**
     * This API allows you to delete the DVM cache <p><strong>OperationId:</strong>clearDVMCache</p>
     * Clear DVM Cache
     */
    async clearDVMCacheRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/services/datavaluemapping/cache`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.StatusFromJSON)(jsonValue));
    }
    /**
     * This API allows you to delete the DVM cache <p><strong>OperationId:</strong>clearDVMCache</p>
     * Clear DVM Cache
     */
    async clearDVMCache(requestParameters, initOverrides) {
        const response = await this.clearDVMCacheRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * This API allows you to delete the default values cache <p><strong>OperationId:</strong>clearDefaultValueCache</p>
     * Clear Default Value Cache
     */
    async clearDefaultValueCacheRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/services/datavaluemapping/defaultValueCache`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.StatusFromJSON)(jsonValue));
    }
    /**
     * This API allows you to delete the default values cache <p><strong>OperationId:</strong>clearDefaultValueCache</p>
     * Clear Default Value Cache
     */
    async clearDefaultValueCache(requestParameters, initOverrides) {
        const response = await this.clearDefaultValueCacheRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * This API allows you to get the converted values <p><strong>OperationId:</strong>getConvertedValues</p>
     * Fetch Converted Values
     */
    async getConvertedValuesRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        if (requestParameters.externalHotelId !== undefined) {
            queryParameters['externalHotelId'] = requestParameters.externalHotelId;
        }
        if (requestParameters.conversionCode) {
            queryParameters['conversionCode'] = requestParameters.conversionCode;
        }
        if (requestParameters.valueToBeConverted) {
            queryParameters['valueToBeConverted'] = requestParameters.valueToBeConverted;
        }
        if (requestParameters.masterValue) {
            queryParameters['masterValue'] = requestParameters.masterValue;
        }
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/hotels/{operaHotelId}/convertedValues`.replace(`{${"operaHotelId"}}`, encodeURIComponent(String(requestParameters.operaHotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.ConvertedValuesFromJSON)(jsonValue));
    }
    /**
     * This API allows you to get the converted values <p><strong>OperationId:</strong>getConvertedValues</p>
     * Fetch Converted Values
     */
    async getConvertedValues(requestParameters, initOverrides) {
        const response = await this.getConvertedValuesRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * This API allows you to get default values <p><strong>OperationId:</strong>getDefaultValues</p>
     * Fetch Default values
     */
    async getDefaultValuesRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/hotels/{operaHotelId}/defaultValues`.replace(`{${"operaHotelId"}}`, encodeURIComponent(String(requestParameters.operaHotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.DefaultValuesFromJSON)(jsonValue));
    }
    /**
     * This API allows you to get default values <p><strong>OperationId:</strong>getDefaultValues</p>
     * Fetch Default values
     */
    async getDefaultValues(requestParameters, initOverrides) {
        const response = await this.getDefaultValuesRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * This API allows you to get the converted hotel code <p><strong>OperationId:</strong>getOperaHotelCode</p>
     * Fetch Converted Hotel Code
     */
    async getOperaHotelCodeRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        if (requestParameters.operaHotelCode !== undefined) {
            queryParameters['operaHotelCode'] = requestParameters.operaHotelCode;
        }
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/hotels/{externalHotelId}/conversions`.replace(`{${"externalHotelId"}}`, encodeURIComponent(String(requestParameters.externalHotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.ConvertedHotelCodeFromJSON)(jsonValue));
    }
    /**
     * This API allows you to get the converted hotel code <p><strong>OperationId:</strong>getOperaHotelCode</p>
     * Fetch Converted Hotel Code
     */
    async getOperaHotelCode(requestParameters, initOverrides) {
        const response = await this.getOperaHotelCodeRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * ping Data Value Mapping Service <p><strong>OperationId:</strong>pingDataValueMappingService</p>
     * ping Data Value Mapping Service
     */
    async pingDataValueMappingServiceRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/services/datavaluemapping/status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.OperaVersionFromJSON)(jsonValue));
    }
    /**
     * ping Data Value Mapping Service <p><strong>OperationId:</strong>pingDataValueMappingService</p>
     * ping Data Value Mapping Service
     */
    async pingDataValueMappingService(requestParameters, initOverrides) {
        const response = await this.pingDataValueMappingServiceRaw(requestParameters, initOverrides);
        return await response.value();
    }
}
exports.DataValueMappingApi = DataValueMappingApi;
/**
 * @export
 */
exports.GetConvertedValuesConversionCodeEnum = {
    AccountType: 'AccountType',
    ActionCode: 'ActionCode',
    ActionType: 'ActionType',
    ActivityLocation: 'ActivityLocation',
    ActivityStatus: 'ActivityStatus',
    ActivityType: 'ActivityType',
    ActivityTypeSc: 'ActivityTypeSc',
    AddressTypes: 'AddressTypes',
    ArrangementCode: 'ArrangementCode',
    BlockCancellationCodes: 'BlockCancellationCodes',
    BlockCodes: 'BlockCodes',
    BlockConversion: 'BlockConversion',
    BlockRanking: 'BlockRanking',
    BlockRateOverrideReason: 'BlockRateOverrideReason',
    BookingStatus: 'BookingStatus',
    BookingType: 'BookingType',
    BusiinessSegment: 'BusiinessSegment',
    CalendarEvent: 'CalendarEvent',
    CancellationCodes: 'CancellationCodes',
    CategoryCode: 'CategoryCode',
    CatItemrate: 'CatItemrate',
    CatRate: 'CatRate',
    CatServing: 'CatServing',
    CatSetup: 'CatSetup',
    ChannelCodes: 'ChannelCodes',
    ComboRoom: 'ComboRoom',
    CommentType: 'CommentType',
    CompetitionCode: 'CompetitionCode',
    CountryCode: 'CountryCode',
    CurrencyCode: 'CurrencyCode',
    DayType: 'DayType',
    DepartmentId: 'DepartmentId',
    DeptNoteCode: 'DeptNoteCode',
    DiscountReason: 'DiscountReason',
    DocumentType: 'DocumentType',
    EventType: 'EventType',
    FitContractType: 'FitContractType',
    Frequency: 'Frequency',
    FunctionSpaceLocation: 'FunctionSpaceLocation',
    FunctionSpaceRateType: 'FunctionSpaceRateType',
    GenderMf: 'GenderMf',
    GtdReq: 'GtdReq',
    GuestPreferenceCode: 'GuestPreferenceCode',
    GuestPreferenceType: 'GuestPreferenceType',
    IndustryCode: 'IndustryCode',
    InfluenceCode: 'InfluenceCode',
    LanguageCodes: 'LanguageCodes',
    MarketingCity: 'MarketingCity',
    MarketingRegion: 'MarketingRegion',
    MarketCode: 'MarketCode',
    MeetingRoomtype: 'MeetingRoomtype',
    MembershipLevel: 'MembershipLevel',
    MembershipType: 'MembershipType',
    MembershipTypeAirline: 'MembershipTypeAirline',
    Nationality: 'Nationality',
    OcrmBounceType: 'OcrmBounceType',
    OcrmEventType: 'OcrmEventType',
    OtaGuestPreferenceCode: 'OtaGuestPreferenceCode',
    PaymentMethod: 'PaymentMethod',
    PhoneType: 'PhoneType',
    PosFamilyGroup: 'PosFamilyGroup',
    PosMajorGroup: 'PosMajorGroup',
    PosMenuDef: 'PosMenuDef',
    PosRevenueCenter: 'PosRevenueCenter',
    PriceCode: 'PriceCode',
    ProductCode: 'ProductCode',
    ProfilePriority: 'ProfilePriority',
    ProfileSource: 'ProfileSource',
    ProfileType: 'ProfileType',
    RateCategory: 'RateCategory',
    RateCode: 'RateCode',
    RateProgram: 'RateProgram',
    RateToRateBlock: 'RateToRateBlock',
    RateType: 'RateType',
    ReasonCode: 'ReasonCode',
    Relationship: 'Relationship',
    ReservationStatus: 'ReservationStatus',
    ReservationType: 'ReservationType',
    ResvBookingMethod: 'ResvBookingMethod',
    RevenueGroups: 'RevenueGroups',
    RevenueType: 'RevenueType',
    Room: 'Room',
    RoomsPotential: 'RoomsPotential',
    RoomCategoryLabel: 'RoomCategoryLabel',
    RoomClass: 'RoomClass',
    RoomPool: 'RoomPool',
    RoomRepairsReasonCode: 'RoomRepairsReasonCode',
    RoomStatusReasons: 'RoomStatusReasons',
    RoutingCode: 'RoutingCode',
    Scope: 'Scope',
    ScopeCity: 'ScopeCity',
    ScAccPriority: 'ScAccPriority',
    ScAccRoomsPot: 'ScAccRoomsPot',
    ScAccSource: 'ScAccSource',
    ScDestination: 'ScDestination',
    ScFunctionSpace: 'ScFunctionSpace',
    ScFunctionSpaceRateCode: 'ScFunctionSpaceRateCode',
    ScFunctionSpaceSetup: 'ScFunctionSpaceSetup',
    ScRevenueTypes: 'ScRevenueTypes',
    ScSetupCode: 'ScSetupCode',
    ScTaskCode: 'ScTaskCode',
    ScTerritory: 'ScTerritory',
    ServiceRequestType: 'ServiceRequestType',
    SgiLanguage: 'SgiLanguage',
    SgiRequestCodes: 'SgiRequestCodes',
    SgiTitle: 'SgiTitle',
    SourceCode: 'SourceCode',
    State: 'State',
    Title: 'Title',
    TraceCode: 'TraceCode',
    TraceDepartment: 'TraceDepartment',
    TransportTypes: 'TransportTypes',
    TravelAgentCommissionCodes: 'TravelAgentCommissionCodes',
    TravelAgentCurrencyCodes: 'TravelAgentCurrencyCodes',
    Trxgenerates: 'Trxgenerates',
    TrxCode: 'TrxCode',
    VipLevel: 'VipLevel',
    WaitlistCodes: 'WaitlistCodes',
    WebuserSecurityQuestion: 'WebuserSecurityQuestion'
};
