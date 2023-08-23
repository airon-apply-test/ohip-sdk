/**
 * OPERA Cloud Content Service
 * Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 */
export declare const ImageCodeType: {
    readonly LogoChain: "LogoChain";
    readonly Logo: "Logo";
    readonly Profile: "Profile";
    readonly Nationality: "Nationality";
    readonly User: "User";
    readonly ToolbarItemCode: "ToolbarItemCode";
    readonly Chain: "Chain";
    readonly MembershipType: "MembershipType";
    readonly MembershipLevel: "MembershipLevel";
    readonly Resort: "Resort";
    readonly RoomTypeChain: "RoomTypeChain";
    readonly RoomType: "RoomType";
    readonly RoomNumber: "RoomNumber";
    readonly OooReasonChain: "OooReasonChain";
    readonly OooReason: "OooReason";
    readonly RoomMaintChain: "RoomMaintChain";
    readonly RoomMaintenance: "RoomMaintenance";
    readonly RoomMaintIssue: "RoomMaintIssue";
    readonly FloorPlan: "FloorPlan";
    readonly SiteLayout: "SiteLayout";
    readonly Attraction: "Attraction";
    readonly PropertyMap: "PropertyMap";
    readonly InvItemChain: "InvItemChain";
    readonly InvItem: "InvItem";
    readonly PackageChain: "PackageChain";
    readonly Package: "Package";
    readonly Pool: "Pool";
    readonly RoomConditionChain: "RoomConditionChain";
    readonly RoomCondition: "RoomCondition";
    readonly FunctionSpace: "FunctionSpace";
    readonly FunctionSetupStyle: "FunctionSetupStyle";
    readonly TrackItTypeChain: "TrackItTypeChain";
    readonly TrackItType: "TrackItType";
    readonly Height: "Height";
    readonly Width: "Width";
};
export type ImageCodeType = typeof ImageCodeType[keyof typeof ImageCodeType];
export declare function ImageCodeTypeFromJSON(json: any): ImageCodeType;
export declare function ImageCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ImageCodeType;
export declare function ImageCodeTypeToJSON(value?: ImageCodeType | null): any;
