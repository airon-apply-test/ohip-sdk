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
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataValueMappingCodeTypeToJSON = exports.DataValueMappingCodeTypeFromJSONTyped = exports.DataValueMappingCodeTypeFromJSON = exports.DataValueMappingCodeType = void 0;
/**
 * Enumeration of the different conversion codes needed for data value mapping.
 * @export
 */
exports.DataValueMappingCodeType = {
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
function DataValueMappingCodeTypeFromJSON(json) {
    return DataValueMappingCodeTypeFromJSONTyped(json, false);
}
exports.DataValueMappingCodeTypeFromJSON = DataValueMappingCodeTypeFromJSON;
function DataValueMappingCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.DataValueMappingCodeTypeFromJSONTyped = DataValueMappingCodeTypeFromJSONTyped;
function DataValueMappingCodeTypeToJSON(value) {
    return value;
}
exports.DataValueMappingCodeTypeToJSON = DataValueMappingCodeTypeToJSON;
